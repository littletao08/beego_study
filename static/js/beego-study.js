/**
 * Created by zhanglida on 15/11/24.
 */
$(function () {

    //边栏滚动控制
    sitebarScrollControll();

    //给赞助添加事件
    $("#sponsor").click(function(){
        $.post("/sponsors/new",function(data){
            $(data).submit();
        })
    });

})


function sitebarScrollControll(){
    var sidebarOuterHeight = $("#sidebar").outerHeight();
    $("div.span9").css("min-height", sidebarOuterHeight);
    var footerOffsetTop = $("footer").position().top;
    var containerOuterHeight = $("header").outerHeight();
    $(window).scroll(function () {

        var $sidebar = $("#sidebar");
        var scrollTop = $(window).scrollTop();
        var headerHeight = $("header").outerHeight() - 10;
        var hasAffix = $sidebar.hasClass("affix");
        var hasAffixTop = $sidebar.hasClass("affix-top");
        var hasAffixBottom = $sidebar.hasClass("affix-bottom");
        //**************设置左边栏固定*****************//
        if (scrollTop >= headerHeight) {
            //**************设置底边栏固定*****************//
            var footerRealTop = footerOffsetTop - scrollTop;
            var compareNumber = (footerOffsetTop - containerOuterHeight);
            if (footerRealTop <= compareNumber && hasAffix) {
                $sidebar.removeClass("affix").addClass("affix-top")
            }
            else if (hasAffixTop && footerRealTop > compareNumber) {
                $sidebar.removeClass("affix-top").addClass("affix")
            }
            /*if(hasAffixTop) {
             $sidebar.removeClass("affix-top").addClass("affix")
             }*/
        }
        else if (scrollTop < headerHeight && (hasAffix || hasAffixBottom)) {
            $sidebar.removeClass("affix").addClass("affix-top")
        }
    });
}
