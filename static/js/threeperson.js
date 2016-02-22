/**
 * Created by zhanglida on 15/11/24.
 */
$(function () {

    var $window = $(window)

    //给赞助添加事件
    $("#sponsor").click(function () {
        $.post("/sponsors/new", function (data) {
            $(data).submit();
        })
    });
    // Disable certain links in docs
    $('section [href^=#]').click(function (e) {
        e.preventDefault()
    })
    //左边栏滚动控制
    setTimeout(function () {
        $('.bs-docs-sidenav').affix({
            offset: {
                top: 270,
                bottom: 270
            }
        })
    }, 100)

    //设置内容区最小告诉和左边栏的高度一致
    var sidebarHeight = $("div#container-left").height();
    $("div#container-right").css("min-height", sidebarHeight);

    //左边脸菜单css切换
    $("#sidebar li").click(function () {
        $($(this).siblings()).removeClass("active")
        $(this).addClass("active")
    });

    //首页点赞
    bindLikeBtnEvent();

    //显示文章内容
    showArticleContent();


});

$(window).load(function(){
    $("div[class$='article-content']").find("pre").addClass("prettyprint linenums");
    prettyPrint();
})

function bindLikeBtnEvent() {
    $(".article-like").click(function () {
        var $like = $(this);
        var articleId = $like.attr("article-id");
        $.ajax({
            url: "/articles/" + articleId + "/likes",
            type: 'post',
            success: function (data) {
                var $likeCount = $like.next();
                if (data.Success) {
                    var oldCount = Number($likeCount.html());
                    if (data.Message > 0) {
                        $likeCount.text(oldCount + 1);
                        $like.attr("class", "glyphicon-heart article-like");
                    } else {
                        $likeCount.text(oldCount - 1);
                        $like.attr("class", "glyphicon-heart-empty article-like");
                    }
                }
                else if (data.Code == 1000) {
                    $("#header-tip").showWarnTip("您未登录");
                }
                else {
                    $("#header-tip").showWarnTip("点赞失败");
                }
            },
            error: function () {
                $("#header-tip").showWarnTip("点赞失败");
            }
        })


    });
}


function initMarkdownEditor(content) {
    var contentError = "<span class=\"help-block form-error\">内容不能为空</span>";

    $("#editor").markdownEditor(
        {
            preview: true,
            onPreview: function (content, callback) {
                callback(marked(content));
            }
        }
    );

    //articleContent
    $('#editor').markdownEditor("setContent", content);

    $("textarea").attr("name", "content").keyup(function () {
        var content = $('#editor').markdownEditor('content');
        if (null != content && content.length > 0) {
            removeContentError();
        } else {
            showContentError();
        }
    });

    $("form[name='article-editor-form']").find("input[name='commit']").click(function () {
        var content = $('#editor').markdownEditor('content');
        if (!$.trim(content)) {
            showContentError();
            return false;
        }
        $("textarea").val(content);
    })


    var showContentError = function () {
        $(".md-editor").css("border-color", "red");
        var $contentError = $(".md-editor").next($(".form-error"));
        if (!$contentError || !$contentError.html()) {
            $(".md-editor").after($(contentError));
        }
    }

    var removeContentError = function () {
        $(".md-editor").css("border-color", "#d8d8d8").next(".form-error").remove();
    }
}


var TipType = {SUCCESS: "alert-success", WARN: "alert-danger", ERROR: "alert-error"};

jQuery.fn.extend({
    showTip: function (message, tipType, callback) {

        var call = function () {
            if (null != callback && $.isFunction(callback)) {
                callback();
            }
            return true;
        }

        $(this).find("span").html(message);
        //从新定义margin-top
        var topBarHeight = $("#topbar").outerHeight();
        var $close = $(this).find("a:first");

        if (tipType == TipType.SUCCESS) {
            $close.css("display", "none");
            $(this).css("margin-top", topBarHeight).addClass(TipType.SUCCESS).removeClass(TipType.ERROR).fadeIn(1000).delay(1000).fadeOut(500, call);
        }

        else if (tipType == TipType.ERROR) {
            $(this).css("margin-top", topBarHeight).addClass(TipType.ERROR).removeClass(TipType.SUCCESS).fadeIn(1000);
            $close.css("display", "block");
            $close.click(function () {
                $("#header-tip").fadeOut(500);
                $close.css("display", "none");
            });
        }
        else if (tipType == TipType.WARN) {
            $close.css("display", "none");
            $(this).css("margin-top", topBarHeight).addClass(TipType.WARN).removeClass(TipType.WARN).fadeIn(1000).delay(1000).fadeOut(500, call);
        }

    },
    showSuccessTip: function (message, callback) {
        $(this).showTip(message, TipType.SUCCESS, callback);
    },
    showErrorTip: function (message, callback) {
        $(this).showTip(message, TipType.ERROR, callback);
    },
    showWarnTip: function (message, callback) {
        $(this).showTip(message, TipType.WARN, callback);
    }

})


function showArticleContent() {
    var $articleContent = $("div[class$='article-content']:hidden");
    if ($articleContent) {
        $articleContent.each(function(){
            var content = marked($(this).text());
            if ($(this).hasClass("full-article-content")) {
                var $content=$("<div>"+content+"</div>");
                $content.find("p").each(function(){
                    var pText = $(this).html();
                    pText=pText.replace(new RegExp('\n','gm'),'<p>');
                    $(this).html(pText);
                });

                $articleContent.html($content).removeClass("hidden");
            } else {
                content = $("<div>").html(content).text();
                $(this).html(content).removeClass("hidden");
            }

        });
    }
}

function registerFormValid(){
    $.formUtils.addValidator(
        {
            name: 'name',
            validatorFunction: function (value, $el, config, language, $form) {
                /******校验用户名长度*********/
                var result = $.formUtils.numericRangeCheck(value.length,'4-20')
                if (result[0]!="ok"){
                    this.errorMessage="用户名长度只能在4-20位字符之间"
                    return false;
                }

                /******校验用户名特殊字符*******/
                var reg=/^[0-9-a-zA-Z\u4E00-\u9FA5_]*$/g;

                if(!reg.test(value)){
                    this.errorMessage="用户名只支持汉字,字母,数字,\"-\",\"_\"组合"
                    return false;
                }

                /******校验用户名是否存在******/
                var result = $.ajax({
                    url: "/users/check_user_name",
                    type: "POST",
                    data: {name: value},
                    async: false,
                    success: function (result) {
                        return result.Success
                    }
                });

                result = JSON.parse(result.responseText);
                this.errorMessage=result.Message
                return result.Success;
            },
        }
    );

    $.formUtils.addValidator(
        {
            name: 'mail',
            validatorFunction: function (value, $el, config, language, $form) {
                var result= $.ajax({
                    url: "/users/check_user_mail",
                    type: "POST",
                    data: {mail: value},
                    async: false
                })
                result = JSON.parse(result.responseText);
                this.errorMessage=result.Message
                return result.Success;
            }
        }
    );

    $.formUtils.addValidator(
        {
            name: 'captcha',
            validatorFunction: function (value, $el, config, language, $form) {

                /******校验用户名长度*********/
                var result = $.formUtils.numericRangeCheck(value.length,'6')
                if (result[0]!="ok"){
                    this.errorMessage="验证码长度必须为6位字符"
                    return false;
                }

                return true ;

            }
        }
    );

}


function setCountdown($elem,countdown,defaultVal) {
    if (countdown == 0) {
        $elem.attr("disabled",false);
        $elem.val(defaultVal);
        return ;
    } else {
        $elem.attr("disabled", true);
        $elem.val("重新发送(" + countdown + ")");
        countdown--;
    }
    setTimeout(function() {
        setCountdown($elem,countdown,defaultVal)
    },1000)
}