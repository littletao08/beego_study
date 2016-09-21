/**
 * 用户自定义脚本.
 */
(function(window, Object, undefined) {

/**
 * 敌人(单个小方块)
 */
var Enemy = qc.defineBehaviour('qc.engine.Enemy', qc.Behaviour, function () {
    this.speed = 8;
    this.radius = 30;
}, {
    // fields need to be serialized
});

// Called when the script instance is being loaded.
Enemy.prototype.awake = function () {
    GameManager.instance.GAME_RESTART.add(this.onGameRestart, this);
};

Enemy.prototype.onGameRestart = function () {
    this.gameObject.destroy();
}

Enemy.prototype.update = function () {
    if (!GameManager.instance.running) {
        return;
    }

    //加分
    if (this.gameObject.anchoredX < 0 && this.gameObject.anchoredX + this.speed >= 0) {
        ScoreManager.instance.add(1);
    }
    this.gameObject.anchoredX += this.speed;

    var xPower = Math.pow(this.gameObject.anchoredX - Player.instance.gameObject.anchoredX, 2);
    var yPower = Math.pow(this.gameObject.anchoredY - Player.instance.gameObject.anchoredY, 2);
    var dist = Math.sqrt(xPower + yPower);

    //判断碰撞
    if (dist < this.radius + Player.instance.radius) {
        console.log("touch");
        GameManager.instance.gameOver();
    }

    //障碍位移出屏幕自动销毁
    if (this.gameObject.anchoredX > this.gameObject.parent.width / 2 + 100) {
        this.gameObject.destroy();
    }
};

/**
 * 障碍模型批量构建
 * 根据pattern创建由多个单个方块构建的障碍
 * pattern.timing 横向位置索引
 * pattern.timing 高度位置索引
 */
var EnemyManager = qc.defineBehaviour('qc.engine.EnemyManager', qc.Behaviour, function () {
    this.pattern = [
        {
            timing: [0, 1, 1, 2],
            spacing: [0, 0, 1, 0]
        },
        {
            timing: [0, 0],
            spacing: [0, 1],
        },
        {
            timing: [0, 0,1],
            spacing: [0, 1,0]
        }
        ,
        {
            timing: [0, 0,1,1],
            spacing: [0, 1,0,1]
        }
        ,
        {
            timing: [0, 0,1],
            spacing: [0, 1,1]
        }
        ,
        {
            timing: [0, 1, 1, 2],
            spacing: [1, 0, 1, 1]
        }
        ,
        {
            timing: [0, 1,1,1],
            spacing: [1, 0,1,2]
        },
        {
            timing: [0, 1],
            spacing: [1,0]
        },
        {
            timing: [0, 0,0,1],
            spacing: [0,1,2, 1]
        }
    ];
    this.tick = 0;
    this.nextTick = -1;
}, {
    enemyPrefab: qc.Serializer.PREFAB
});

EnemyManager.prototype.awake = function () {
    GameManager.instance.GAME_START.add(this.onGameStart, this);
    GameManager.instance.GAME_OVER.add(this.onGameOver, this);
    GameManager.instance.GAME_RESTART.add(this.onGameRestart, this);

};

EnemyManager.prototype.update = function () {
    if (!GameManager.instance.running) {
        return;
    }
    this.tick++;
    if (this.tick == this.nextTick) {
        this.spawn();
    }
};

EnemyManager.prototype.onGameStart = function () {
    this.tick = 0;
    this.nextTick = -1;
    this.spawn();
};

EnemyManager.prototype.onGameOver = function () {
    this.onGameStart();
};

EnemyManager.prototype.onGameRestart = function () {
    this.spawn();
};

//根君障碍模型创建障碍
EnemyManager.prototype.spawn = function () {

    var ppt = this.pattern[Math.floor(Math.random() * this.pattern.length)];
    for (var i = 0; i < ppt.timing.length; i++) {
        var enemy = this.game.add.clone(this.enemyPrefab, this.gameObject);
        enemy.anchoredX = -800 - ppt.timing[i] * 100;
        enemy.anchoredY = -100 - ppt.spacing[i] * 100;
    }

    this.nextTick = this.tick + ppt.timing.length * 20 + 100;
}

/**
 * 游戏开始结束控制
 */
var GameManager = qc.defineBehaviour('qc.engine.GameManager', qc.Behaviour, function () {
    GameManager.instance = this;

    this.GAME_START = new qc.Signal();
    this.GAME_OVER = new qc.Signal();
    this.GAME_RESTART = new qc.Signal();
    this.running = false;

}, {});

GameManager.prototype.awake = function () {
    this.gameStart();
};

// Called every frame, if the behaviour is enabled.
//GameManager.prototype.update = function() {
//
//};


GameManager.prototype.gameStart = function () {
    this.GAME_START.dispatch();
    this.running = true;

};


GameManager.prototype.gameOver = function () {
    this.GAME_OVER.dispatch();
    this.running = false;
};


GameManager.prototype.restart = function () {
    this.GAME_RESTART.dispatch();
    this.running = true;
};
/**
 * 游戏结束面板管理
 */
var GameOverManager = qc.defineBehaviour('qc.engine.GameOverManager', qc.Behaviour, function() {
}, {
    retryButton:qc.Serializer.NODE,
    score:qc.Serializer.NODE,
});

GameOverManager.prototype.awake = function() {
    this.gameObject.visible = false ;
    GameManager.instance.GAME_OVER.add(this.onGameOver,this);
    this.retryButton.onClick.add(this.onClickRetry,this);
    this.game.input.onKeyDown.add(this.onKeyDown,this);
};

//游戏结束
GameOverManager.prototype.onGameOver = function() {
    this.gameObject.visible = true ;
    this.score.text = ScoreManager.instance.score+'';
}

//重新开始
GameOverManager.prototype.onClickRetry = function() {
    this.gameObject.visible = false ;
    GameManager.instance.restart();
}

//游戏结束按回车键重新开始
GameOverManager.prototype.onKeyDown = function(keyCode) {
    if (keyCode == 13 && !GameManager.instance.running) {
        this.onClickRetry();
    }
}
/**
 * 游戏主角
 */
var Player = qc.defineBehaviour('qc.engine.Player', qc.Behaviour, function () {
    Player.instance = this;
    this.speed = 0;
    this.gravity = 0.6;
    this.groundHeight = 100;
    this.doubleJumped = false;
    this.radius = 50;

    this.jumpHeight = 20;
    this.doubleJumpHeight = 15;
    this.smashSpeed = 35;

    this.dragStartY = 0;
}, {

});

Player.prototype.awake = function () {
    GameManager.instance.GAME_START.add(this.onGameStart, this);
    GameManager.instance.GAME_OVER.add(this.onGameOver, this);
    GameManager.instance.GAME_RESTART.add(this.onGameRestart, this);
    // 注册点击响应事件
    this.game.input.onPointerDown.add(this.onPointerDown, this);
    //this.game.input.onPointerMove.add(this.onPointerMove,this);
    this.game.input.onPointerUp.add(this.onPointerUp, this);

};
//游戏开始
Player.prototype.onGameStart = function () {
    this.game.input.onKeyDown.add(this.onKeyDown, this);
    this.gameObject.anchoredY = -this.groundHeight;
};

//游戏重新开始
Player.prototype.onGameRestart = function () {
    this.onGameStart();
};

//游戏结束
Player.prototype.onGameOver = function () {
    this.game.input.onKeyDown.remove(this.onKeyDown, this);
};


Player.prototype.update = function () {
    if (!GameManager.instance.running) {
        return;
    }

    this.speed += this.gravity;

    if (this.gameObject.anchoredY < -this.groundHeight && this.gameObject.anchoredY + this.speed >= -this.groundHeight) {
        this.land();
    }
    this.gameObject.anchoredY += this.speed;


    if (this.gameObject.anchoredY > -this.groundHeight) {
        this.gameObject.anchoredY = -this.groundHeight;
        this.doubleJumped = false;
    }
};

//捕获键盘输入
Player.prototype.onKeyDown = function (keyCode) {
    //if(keyCode == qc.Keyboard.UP){
    if (keyCode == 65) {
        this.jump();
        //}else if (keyCode == qc.Keyboard.DOWN){
    } else if (keyCode == 68) {
        this.smash();
    }
};

//跳跃
Player.prototype.jump = function () {
    if (this.gameObject.anchoredY >= -this.groundHeight) {
        this.speed = -this.jumpHeight;
    } else if (!this.doubleJumped) {
        this.speed = -this.doubleJumpHeight;
        this.doubleJumped = true;

    }
};

//下降
Player.prototype.smash = function () {
    this.speed = this.smashSpeed;
};

//着陆
Player.prototype.land = function () {
    this.gameObject.Animator.play("land");
};

//滑动开始
Player.prototype.onPointerDown = function (id, x, y) {
    this.dragStartY = y;
};

//滑动结束
Player.prototype.onPointerUp = function (id, x, y) {
    var diff = y - this.dragStartY ;
    if (diff == 0){
        return ;
    }
    if (y - this.dragStartY < 0 ) {
        this.jump();
    } else if (y - this.dragStartY > 0) {
        this.smash();
    }
};


// define a user behaviour
var ScoreManager = qc.defineBehaviour('qc.engine.ScoreManager', qc.Behaviour, function() {
    ScoreManager.instance = this ;
    this.score = 0 ;
}, {
    // fields need to be serialized
});

ScoreManager.prototype.awake = function() {
    GameManager.instance.GAME_RESTART.add(this.onGameRestart,this);
};

//重新开始
ScoreManager.prototype.onGameRestart = function() {
    this.score = 0 ;
    this.gameObject.text = this.score+'';
};

//增加分数
ScoreManager.prototype.add = function(point) {
    this.score += point;
    this.gameObject.text = this.score+'';
};


}).call(this, this, Object);
