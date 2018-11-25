function addMsg() {
    if($('#name').val() == '') {
        $('#nameTip').show();
        $('#nameTip').text('请填写姓名');
        return false
    } else {
        $('#nameTip').hide();
    }
    if($('#email').val() == '') {
        $('#emailTip').show();
        $('#emailTip').text('请填写邮箱');
        return false
    } else {
        $('#emailTip').hide();
    }
    if($('#content').val() == '') {
        $('#contentTip').show();
        $('#contentTip').text('请填写留言内容');
        return false
    } else {
        $('#contentTip').hide();
    }
    $.ajax({
        type: "POST",
        url: "/l/add",
        data: $('#addForm').serialize(),
        success: function (data) {
            $('#name').val('');
            $('#email').val('');
            $('#content').val('');
            alert("留言成功!")
            initPage();
        },
        error: function(err){
            console.log(err)
        }
    });
    return false;
}

function initPage() {
    $.ajax({
        type: "GET",
        url: "/l/top",
        success: function (data) {
            var html = ""
            $(data).each(function (i, n) {
                console.log(n)
                html += '<div class="fb"><ul><p class="fbtime"><span>' + n.createTime + '</span>' + n.name + '</p><p class="fbinfo">' + n.content + '</p></ul></div>';
                if(n.replyContent) {
                    html += '<div class="hf"><ul><p class="zzhf"><font color="#FF0000">站长回复:</font>' + n.replyContent + '</p></ul></div>'
                }
            });
            $('#gbookDiv').html(html)
        },
        error: function(err){
            console.log(err)
        }
    });
}
initPage();