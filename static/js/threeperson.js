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
    $("div#container-right").css("min-height",sidebarHeight);

    //左边脸菜单css切换
    $("#sidebar li").click(function () {
        $($(this).siblings()).removeClass("active")
        $(this).addClass("active")
    });
    //首页点赞
    bindLikeBtnEvent();

    //首页评论
    $(".article-item-bottom").find(".comment").click(function () {
        $("#header-tip").showWarnTip("评论失败!");
    });


    /*//给文章详情页修改绑定事件
    $(".article-detail-nav-right").find(".glyphicon-pencil").click(function(){
        alert(1);
    });
*/
});

function bindLikeBtnEvent(){
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
                        $like.attr("class","glyphicon-heart article-like");
                    } else {
                        $likeCount.text(oldCount - 1);
                        $like.attr("class","glyphicon-heart-empty article-like");
                    }
                }
                else if(data.Code==1000){
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
    $('#editor').markdownEditor("setContent",content);

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
