$(document).ready(function () {
    //nav
    $("#mnavh").click(function () {
        $("#starlist").toggle();
        $("#mnavh").toggleClass("open");
    });

    var obj = null;
    var As = document.getElementById('starlist').getElementsByTagName('a');
    obj = As[0];
    for (i = 1; i < As.length; i++) {
        if (window.location.href.indexOf(As[i].href) >= 0)
            obj = As[i];
    }
    obj.id = 'selected';

    var new_scroll_position = 0;
    var last_scroll_position;
    var header = document.getElementById("header");

    window.addEventListener('scroll', function (e) {
        last_scroll_position = window.scrollY;

        // Scrolling down
        if (new_scroll_position < last_scroll_position && last_scroll_position > 80) {
            // header.removeClass('slideDown').addClass('slideUp');
            header.classList.remove("slideDown");
            header.classList.add("slideUp");

            // Scrolling up
        } else if (new_scroll_position > last_scroll_position) {
            // header.removeClass('slideUp').addClass('slideDown');
            header.classList.remove("slideUp");
            header.classList.add("slideDown");
        }

        new_scroll_position = last_scroll_position;
    });

    // browser window scroll (in pixels) after which the "back to top" link is shown
    var offset = 300,
        //browser window scroll (in pixels) after which the "back to top" link opacity is reduced
        offset_opacity = 1200,
        //duration of the top scrolling animation (in ms)
        scroll_top_duration = 700,
        //grab the "back to top" link
        $back_to_top = $('.cd-top');

    //hide or show the "back to top" link
    $(window).scroll(function () {
        ($(this).scrollTop() > offset) ? $back_to_top.addClass('cd-is-visible') : $back_to_top.removeClass('cd-is-visible cd-fade-out');
        if ($(this).scrollTop() > offset_opacity) {
            $back_to_top.addClass('cd-fade-out');
        }
    });
    //smooth scroll to top
    $back_to_top.on('click', function (event) {
        event.preventDefault();
        $('body,html').animate({
            scrollTop: 0,
        }, scroll_top_duration
        );
    });

    //aside
    var Sticky = new hcSticky('aside', {
        stickTo: 'main',
        innerTop: 200,
        followScroll: false,
        queries: {
            480: {
                disable: true,
                stickTo: 'body'
            }
        }
    });

    $.ajax({
        type: "GET",
        url: "/meun/all",
        success: function (data) {
            //console.log(data)
        },
        error: function(err){
            console.log(err)
        }
    });

    $.ajax({
        type: "GET",
        url: "/my/all",
        success: function (data) {
            $('#myDesc').html("<b>" + data[0].Name + "</b>" + data[0].MyDesc)
        },
        error: function(err){
            console.log(err)
        }
    });

    $.ajax({
        type: "GET",
        url: "/p/top",
        success: function (data) {
            var html = ""
            $(data).each(function (i, n) {
                html += '<li><a href="/p/all"><img src="' + n.url + '"></a></li>'
            })
            $('#mainPicture').html(html)
        },
        error: function(err){
            console.log(err)
        }
    });

    $.ajax({
        type: "GET",
        url: "/c/main",
        success: function (data) {
            var html = ""
            $(data).each(function (i, n) {
                html += '<li><a href="/p/all">' + n.name + '(' + n.count + ')' + '</a></li>'
            })
            $('#mainCategory').html(html)
        },
        error: function(err){
            console.log(err)
        }
    });

    $.ajax({
        type: "GET",
        url: "/a/main",
        success: function (data) {
            var html = ""
            $(data).each(function (i, n) {
                var inhtml = '<li><i><a href="/"><img src="images/1.jpg"></a></i>\n' +
                    '                <h3><a href="/">' + n.title + '</a></h3>\n' +
                    '            <p>' + n.introduction + '</p>\n' +
                    '            </li>';
                html += inhtml
            });
            $('#mainContent').html(html);
        },
        error: function(err){
            console.log(err)
        }
    });

    //站长推荐及最高阅读量
    $.ajax({
        type: "GET",
        url: "/a/readtop",
        success: function (data) {
            var html = ""
            $(data).each(function (i, n) {
                html += "<li><a href=\"/\">" + n.title +  "</a></li>"
            });
            $('#topRead').html(html)
        },
        error: function(err){
            console.log(err)
        }
    });

    $.ajax({
        type: "GET",
        url: "/l/top",
        success: function (data) {
            var html = ""
            $(data).each(function (i, n) {
                html += "<div class="fb"><ul><p class="fbtime"><span> 2018-07-21</span> 夜月归途</p><p class="fbinfo">从青姐朋友圈分享的文章《我为什么要做个人网站》过来的，自习看了下你的网站非常不错，看的出来你一直在坚持!</p></ul></div><div class="hf"><ul><p class="zzhf"><font color="#FF0000">站长回复:</font>感谢捧场啊！看了你的网站，有两个月没更新了哦~</p></ul></div>"
            });
            $('#topRead').html(html)
        },
        error: function(err){
            console.log(err)
        }
    });
});