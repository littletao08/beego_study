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
    //左边栏滚动控制
    setTimeout(function () {
        $('.bs-docs-sidenav').affix({
            offset: {
                top: function () {
                    return 270
                }
                ,
                bottom: 270
            }
        })
    }, 100)

    //左边脸菜单css切换
    $("#sidebar li").click(function () {
        $($(this).siblings()).removeClass("active")
        $(this).addClass("active")
    });
    //首页点赞
    $("div[class^='article-detail-nav']").find(".like").click(function () {
        //后台发送请求,文章点赞+1 ,同一个用户只能有一次点赞;
        //后台请求返回点赞成功;返回累计点赞的数量
        $.ajax("/article/praise/2", {
            type: 'post',
            success: function (data) {
                $("#header-tip").showSuccessTip("点赞成功!");
                $(this).find("i").html("(1)");
            }, error: function (data) {
                $("#header-tip").showSuccessTip(data);
            }
        })



    });

    //首页评论
    $(".article-item-bottom").find(".comment").click(function () {
        $("#header-tip").showErrorTip("评论失败!");
    });

})


function initMarkdownEditor() {

    var contentError = "<span class=\"help-block form-error\">内容不能为空</span>";

    $("#editor").markdownEditor(
        {
            preview: true,
            onPreview: function (content, callback) {
                callback(marked(content));
            }
        }
    );
    $("textarea").attr("name", "content").keyup(function () {
        var content = $('#editor').markdownEditor('content');
        if (null != content && content.length > 0) {
            removeContentError();
        } else {
            showContentError();
        }
    });

    $("form#articleEditorForm").find("input[name='commit']").click(function () {
        var content = $('#editor').markdownEditor('content');
        if (!$.trim(content)) {
            showContentError();
            return false;
        }

        $("textarea").val(content);
    })


    var showContentError = function () {
        $(".md-editor").css("border-color", "red");
        //.after($(contentError))
        var $contentError = $(".md-editor").next($(".form-error"));
        if (!$contentError || !$contentError.html()) {
            $(".md-editor").after($(contentError));
        }
    }

    var removeContentError = function () {
        $(".md-editor").css("border-color", "#d8d8d8").next().remove();
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
