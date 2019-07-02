/*!zdialog | (c) 2016 | by zcy*/ (function($, window, document, undefined) {
    var PARAMS;
    var DEFAULTPARAMS = {
        Title:"标题",
        Content:"",
        BtnL:"确定",
        BtnR:"取消",
        FunL:new Object(),
        FunR:new Object()
    };
    $.DialogByZ = {
        Alert:function(params) {
            Show(params, "Alert");
        },
        Confirm:function(params) {
            Show(params, "Confirm");
        },
        Autofade:function(params) {
            Show(params, "Autofade");
        },
        Close:function() {
            $(".zbox-popup,.zbox-popup-backdrop").remove();
        },
        Loading:function(Url) {
            loadBox(Url);
        }
    };
    function Init(params) {
        if (params != undefined && params != null) {
            PARAMS = $.extend({}, DEFAULTPARAMS, params);
        }
    }
    function loadBox(Url) {
        var url = Url;
        var dislogContainer = $('<div class="zbox-popup" style="display: block;"><img  id="zchange" src="' + url + '"></div>');
        var blackFilter = $('<div class="zbox-popup-backdrop" style="display: block;"></div>');
        setTimeout(function() {
            $(".zbox-popup").addClass("zbox-popup-in");
            $(".zbox-popup-backdrop").addClass("zbox-active");
        }, 30);
        $("body").append(dislogContainer);
    }
    function Show(params, caller) {
        Init(params);
        var dislogContainer;
        var dialogInner;
        var dialogBtn;
        var blackFilter = $('<div class="zbox-popup-backdrop" style="display: block;"></div>');
        if (caller == "Autofade") {
            dislogContainer = $('<div class="zbox-toast-container"><div class="zbox-toast-message">' + PARAMS.Content + "</div></div>");
            $("body").append(dislogContainer);
            setTimeout(function() {
                $(".zbox-toast-container").addClass("zbox-active");
            }, 30);
            setTimeout(function() {
                $(".zbox-toast-container").remove();
            }, 3e3);
        } else {
            dislogContainer = $('<div class="zbox-popup" style="display: block;"></div>');
            dialogInner = $('<div class="zbox-popup-inner"><div class="zbox-popup-title">' + PARAMS.Title + '</div><div class="zbox-popup-text">' + PARAMS.Content + "</div></div>");
            dialogBtn = $('<div class="zbox-popup-buttons"><span class="zbox-popup-button" index="0">' + PARAMS.BtnL + "</span></div>");
            if (caller == "Confirm") {
                dialogBtn.append($('<span class="zbox-popup-button R" index="1">' + PARAMS.BtnR + "</span>"));
            }
            dislogContainer.append(dialogInner);
            dislogContainer.append(dialogBtn);
            setTimeout(function() {
                $(".zbox-popup").addClass("zbox-popup-in");
                $(".zbox-popup-backdrop").addClass("zbox-active");
            }, 10);
            $("body").append(blackFilter);
            $("body").append(dislogContainer);
            $(".zbox-popup-button").click(function() {
                var indexs = $(this).attr("index");
                if (indexs == 0) {
                    if ($.isFunction(PARAMS.FunL)) {
                        PARAMS.FunL();
                    } else {
                        $.DialogByZ.Close();
                    }
                } else {
                    if ($.isFunction(PARAMS.FunR)) {
                        PARAMS.FunR();
                    } else {
                        $.DialogByZ.Close();
                    }
                }
                return false;
            });
        }
    }
})(jQuery, window, document);