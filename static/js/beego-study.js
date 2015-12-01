
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
                    return $window.width() <= 980 ? 320 : 230
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
    $(".article-item-bottom").find(".praise").click(function(){
        $("#header-tip").showSuccessTip("点赞成功!");
    });


})


function initMarkdownEditor() {

    $("#editor").markdownEditor(
        {
            preview: true,
            onPreview: function (content, callback) {
                callback(marked(content));
            },
            onSave: function (content) {
                var content = marked(content);
                if (!$.trim(content)) {
                    $(".md-editor").css("border", "1px solid red");
                    return;
                }
                $.post('/articles', {title: "go study", content: content},
                    function (data) {
                        if (data) {
                            $("#header-tip").showSuccessTip("保存成功!",function(){
                                window.location.href = "/"
                            });
                        } else {
                            $("#header-tip").showSuccessTip("保存失败!")
                        }
                    }
                )
            }
        }
    );

    $("textarea").keyup(function () {
        var content = $('#editor').markdownEditor('content');
        if (null != content && content.length > 0) {
            $(".md-editor").css("border", "");
        } else {
            $(".md-editor").css("border", "1px solid red");
        }
    });


}



var TipType = {SUCCESS: "alert-success", WARN: "warn", ERROR: "alert-error"};

jQuery.fn.extend({
    showTip: function (message, tipType, callback) {
        var call = function(){
            if (null != callback && $.isFunction(callback)) {
                callback();
            }
            return true ;
        }
        if (tipType == TipType.SUCCESS) {
            $(this).find("span").html(message);
            $(this).addClass(TipType.SUCCESS).removeClass(TipType.ERROR).fadeIn(1000).delay(1000).fadeOut(500,call);
        }
        else if (tipType == TipType.ERROR) {
            $(this).find("span").html(message);
            $(this).addClass(TipType.ERROR).removeClass(TipType.SUCCESS).fadeIn(1000).delay(1000,call)
        }

    },
    showSuccessTip: function (message, callback) {
        $(this).showTip(message, TipType.SUCCESS, callback);
    },
    showErrorTip: function (message, callback) {
        $(this).showTip(message, TipType.ERROR, callback);
    }

})
