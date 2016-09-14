/**
 * 用户自定义脚本.
 */
(function(window, Object, undefined) {

// define a user behaviour
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

// Called every frame, if the behaviour is enabled.
Enemy.prototype.update = function () {
    if (!GameManager.instance.running) {
        return;
    }

    if (this.gameObject.anchoredX < 0 && this.gameObject.anchoredX + this.speed >= 0) {
        ScoreManager.instance.add(1);
    }
    this.gameObject.anchoredX += this.speed;

    var xPower = Math.pow(this.gameObject.anchoredX - Player.instance.gameObject.anchoredX, 2);
    var yPower = Math.pow(this.gameObject.anchoredY - Player.instance.gameObject.anchoredY, 2);

    var dist = Math.sqrt(xPower + yPower);

    if (dist < this.radius + Player.instance.radius) {
        console.log("touch");
        GameManager.instance.gameOver();
    }
    if (this.gameObject.anchoredX > this.gameObject.parent.width / 2 + 100) {
        this.gameObject.destroy();
    }
};

var EnemyManager = qc.defineBehaviour('qc.engine.EnemyManager', qc.Behaviour, function () {
    this.pattern = [
        {
            timing: [0, 1, 1, 2],
            spacing: [0, 0, 1, 0],
            leng: 3
        },
        {
            timing: [0, 0],
            spacing: [0, 1],
            leng: 1
        },
        {
            timing: [0, 0,1],
            spacing: [0, 1,0],
            leng: 2
        }
        ,
        {
            timing: [0, 0,1,1],
            spacing: [0, 1,0,1],
            leng: 2
        }
        ,
        {
            timing: [0, 0,1],
            spacing: [0, 1,1],
            leng: 2
        }
        ,
        {
            timing: [0, 1, 1, 2],
            spacing: [1, 0, 1, 1],
            leng: 3
        }
        ,
        {
            timing: [0, 1,1,1],
            spacing: [1, 0,1,2],
            leng: 2
        },
        {
            timing: [0, 1],
            spacing: [1,0],
            leng: 2
        },
        {
            timing: [0, 0,0,1],
            spacing: [0,1,2, 1],
            leng: 2
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

EnemyManager.prototype.spawn = function () {

    var ppt = this.pattern[Math.floor(Math.random() * this.pattern.length)];
    for (var i = 0; i < ppt.timing.length; i++) {
        var enemy = this.game.add.clone(this.enemyPrefab, this.gameObject);
        enemy.anchoredX = -800 - ppt.timing[i] * 100;
        enemy.anchoredY = -100 - ppt.spacing[i] * 100;
    }

    this.nextTick = this.tick + ppt.leng * 20 + 100;
}

// define a user behaviour
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

GameOverManager.prototype.onGameOver = function() {
    this.gameObject.visible = true ;
    this.score.text = ScoreManager.instance.score+'';
}


GameOverManager.prototype.onClickRetry = function() {
    this.gameObject.visible = false ;
    GameManager.instance.restart();
}

GameOverManager.prototype.onKeyDown = function(keyCode) {
    if (keyCode == 13 && !GameManager.instance.running) {
        this.onClickRetry();
    }
}
// define a user behaviour
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
    // fields need to be serialized

});

// Called when the script instance is being loaded.
Player.prototype.awake = function () {
    GameManager.instance.GAME_START.add(this.onGameStart, this);
    GameManager.instance.GAME_OVER.add(this.onGameOver, this);
    GameManager.instance.GAME_RESTART.add(this.onGameRestart, this);
    // 注册点击响应事件
    this.game.input.onPointerDown.add(this.onPointerDown, this);
    //this.game.input.onPointerMove.add(this.onPointerMove,this);
    this.game.input.onPointerUp.add(this.onPointerUp, this);

};

Player.prototype.onGameStart = function () {
    this.game.input.onKeyDown.add(this.onKeyDown, this);
    this.gameObject.anchoredY = -this.groundHeight;
};

Player.prototype.onGameRestart = function () {
    this.onGameStart();
};

Player.prototype.onGameOver = function () {
    this.game.input.onKeyDown.remove(this.onKeyDown, this);
};

// Called every frame, if the behaviour is enabled.
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

Player.prototype.onKeyDown = function (keyCode) {
    //if(keyCode == qc.Keyboard.UP){
    if (keyCode == 65) {
        this.jump();
        //}else if (keyCode == qc.Keyboard.DOWN){
    } else if (keyCode == 68) {
        this.smash();
    }
};

Player.prototype.jump = function () {
    if (this.gameObject.anchoredY >= -this.groundHeight) {
        this.speed = -this.jumpHeight;
    } else if (!this.doubleJumped) {
        this.speed = -this.doubleJumpHeight;
        this.doubleJumped = true;

    }
};


Player.prototype.smash = function () {
    this.speed = this.smashSpeed;
};

Player.prototype.land = function () {
    this.gameObject.Animator.play("land");
};


Player.prototype.onPointerDown = function (id, x, y) {
    this.dragStartY = y;
};


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

// Called when the script instance is being loaded.
ScoreManager.prototype.awake = function() {
    GameManager.instance.GAME_RESTART.add(this.onGameRestart,this);
};

ScoreManager.prototype.onGameRestart = function() {
    this.score = 0 ;
    this.gameObject.text = this.score+'';
};

// Called every frame, if the behaviour is enabled.
ScoreManager.prototype.add = function(point) {
    this.score += point;
    this.gameObject.text = this.score+'';
};


}).call(this, this, Object);
