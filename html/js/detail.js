$(document).ready(function () {
    let id = getQueryString('id')
    $('#detailUrl').attr('href', 'detail.html?id=' + id);
    $.ajax({
        type: "POST",
        url: "/a/detail/" + id,
        success: function (data) {
            $('#aTitle').text(data.title);
            $('#aAuthor').text(data.author);
            $('#aCreateTime').text('时间：' + data.createTime);
            $('#aLabel').html(getLabel(data));
            $('#aIntroduction').append(data.introduction);
            $('#aContent').html(data.content);
        },
        error: function(err){
            console.log(err)
        }
    });

    $.ajax({
        type: "GET",
        url: "/e/article/" + id,
        success: function (data) {
            $('#aReadCount').text(data.readCount + '人已阅读');
            $('#aLike').html(getLike(data));
            $('#aPreA').append(getPreAOrNextA(data, 'p'))
            $('#aNextA').append(getPreAOrNextA(data, 'n'))
        },
        error: function(err){
            console.log(err)
        }
    });

    //查询评论
    $.ajax({
        type: "Post",
        url: "/co/list/article/" + id,
        success: function (data) {
            let html = "";
            $(data).each(function (i, n) {
                html += '<div class="fb"><ul><p class="fbtime"><span>' + n.createTime + '</span>' + n.commentator + '</p><p class="fbinfo">' + n.content + '</p></ul></div>';
            });
            $('#commonDiv').prepend(html)
        },
        error: function(err){
            console.log(err)
        }
    });

    //初始化验证码
    genCaptcha(id);

    $('#captchaCode').click(function() {
            genCaptcha(id);
        }
    );
});

function genCaptcha(id) {
    $.ajax({
        type: "get",
        url: "/ca/gen/captcha/" + id,
        success: function (data) {
            $('#captchaCode').attr("src", data);
        },
        error: function(err){
            console.log(err)
        }
    });
}

function getQueryString(name) {
    var reg = new RegExp("(^|&)" + name + "=([^&]*)(&|$)", "i");
    var r = window.location.search.substr(1).match(reg);
    if (r != null) return unescape(r[2]);
    return null;
}

function getLabel(data) {
    let html = "";
    //var labelHtml = '<a href="#" target="_blank">个人博客</a> &nbsp; <a href="#" target="_blank">小世界</a>';
    if(data.label) {
        let label = data.label;
        let ls = label.split(";")
        $(ls).each(function (i, n) {
            if (i != 0) {
                html += '<a href="#" target="_blank">' + n +  '</a>';
            } else {
                html += ' &nbsp;<a href="#" target="_blank">' + n +  '</a>';
            }
        })
    }
    return html;
}

function getLike(data) {
    let html = '<a href="JavaScript:AddLike(\'' + data.articleId + '\');"> 很赞哦！(<b id="diggnum">' + data.likeCount + '</b>)</a>';
    return html;
}

function getPreAOrNextA(data, type) {
    if(type === 'p') {
        if(data.preArticle) {
            return '<a href="detail.html?id=' + data.preArticle + '">' + data.preArticleName + '</a>';
        } else {
            return '<a href="list.html">返回列表</a>';
        }
    } else {
        if(data.nextArticle) {
            return '<a href="detail.html?id=' + data.nextArticle + '">' + data.nextArticleName + '</a>';
        } else {
            return '<a href="list.html">返回列表</a>';
        }
    }
}

function AddLike(aid) {
    $.ajax({
        type: "GET",
        url: "/e/addlike/" + aid,
        success: function (data) {
            let html = '<a href="JavaScript:return false;");"> 已赞！(<b id="diggnum">' + data.likeCount + '</b>)</a>';
            $('#aLike').html(html);
        },
        error: function(err){
            console.log(err)
        }
    });
}