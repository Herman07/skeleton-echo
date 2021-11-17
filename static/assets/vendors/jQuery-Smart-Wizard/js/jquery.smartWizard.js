/*
 * SmartWizard 3.3.1 plugin
 * jQuery Wizard control Plugin
 * by Dipu
 *
 * Refactored and extended:
 * https://github.com/mstratman/jQuery-Smart-Wizard
 *
 * Original URLs:
 * http://www.techlaboratory.net
 * http://tech-laboratory.blogspot.com
 */

function SmartWizard(target, options) {
    this.target       = target;
    this.options      = options;
    this.curStepIdx   = options.selected;
    this.steps        = $(target).children("ul").children("li").children("a"); // Get all anchors
    this.contentWidth = 0;
    this.msgBox = $('<div class="msgBox"><div class="content"></div><a href="#" class="close">X</a></div>');
    this.elmStepContainer = $('<div></div>').addClass("stepContainer");
    this.loader = $('<div>Loading</div>').addClass("loader");
    this.buttons = {
        next : $('<button>'+options.labelNext+'</button>').attr("type","button").addClass("buttonNext"),
        previous : $('<button>'+options.labelPrevious+'</button>').attr("type","button").addClass("buttonPrevious"),
        finish  : $('<button>'+options.labelFinish+'</button>').attr("type","submit").addClass("buttonFinish")
    };

    /*
     * Private functions
     */

    var _init = function($this) {
        var elmActionBar = $('<div></div>').addClass("actionBar");
        elmActionBar.append($this.msgBox);
        $('.close',$this.msgBox).click(function() {
            $this.msgBox.fadeOut("normal");
            return false;
        });

        var allDivs = $this.target.children('div');
        $this.target.children('ul').addClass("anchor");
        allDivs.addClass("content");

        // highlight steps with errors
        if($this.options.errorSteps && $this.options.errorSteps.length>0){
            $.each($this.options.errorSteps, function(i, n){
                $this.setError({ stepnum: n, iserror:true });
            });
        }

        $this.elmStepContainer.append(allDivs);
        elmActionBar.append($this.loader);
        $this.target.append($this.elmStepContainer);
        elmActionBar.append($this.buttons.finish)
                    .append($this.buttons.next)
                    .append($this.buttons.previous);
        $this.target.append(elmActionBar);
        this.contentWidth = $this.elmStepContainer.width();

        $($this.buttons.next).click(function() {
            $this.goForward();
            return false;

        });
        $($this.buttons.previous).click(function() {
            $this.goBackward();
            return false;
        });
        $($this.buttons.finish).click(function() {
            if(!$(this).hasClass('buttonDisabled')){
                if($.isFunction($this.options.onFinish)) {
                    var context = { fromStep: $this.curStepIdx + 1 };
                    if(!$this.options.onFinish.call(this,$($this.steps), context)){
                        return false;
                    }
                }else{
                    var frm = $this.target.parents('form');
                    if(frm && frm.length){
                        frm.submit();
                    }
                }
            }
            return false;
        });

        $($this.steps).bind("click", function(e){
            if($this.steps.index(this) == $this.curStepIdx){
                return false;
            }
            var nextStepIdx = $this.steps.index(this);
            var isDone = $this.steps.eq(nextStepIdx).attr("isDone") - 0;
            if(isDone == 1){
                _loadContent($this, nextStepIdx);
            }
            return false;
        });

        // Enable keyboard navigation
        if($this.options.keyNavigation){
            $(document).keyup(function(e){
                if(e.which==39){ // Right Arrow
                    $this.goForward();
                }else if(e.which==37){ // Left Arrow
                    $this.goBackward();
                }
            });
        }
        //  Prepare the steps
        _prepareSteps($this);
        // Show the first slected step
        _loadContent($this, $this.curStepIdx);
    };

    var _prepareSteps = function($this) {
        if(! $this.options.enableAllSteps){
            $($this.steps, $this.target).removeClass("selected").removeClass("done").addClass("disabled");
            $($this.steps, $this.target).attr("isDone",0);
        }else{
            $($this.steps, $this.target).removeClass("selected").removeClass("disabled").addClass("done");
            $($this.steps, $this.target).attr("isDone",1);
        }

        $($this.steps, $this.target).each(function(i){
            $($(this).attr("href").replace(/^.+#/, '#'), $this.target).hide();
            $(this).attr("rel",i+1);
        });
    };

    var _step = function ($this, selStep) {
        return $(
            $(selStep, $this.target).attr("href").replace(/^.+#/, '#'),
            $this.target
        );
    };

    var _loadContent = function($this, stepIdx) {
        var selStep = $this.steps.eq(stepIdx);
        var ajaxurl = $this.options.contentURL;
        var ajaxurl_data = $this.options.contentURLData;
        var hasContent = selStep.data('hasContent');
        var stepNum = stepIdx+1;
        if (ajaxurl && ajaxurl.length>0) {
            if ($this.options.contentCache && hasContent) {
                _showStep($this, stepIdx);
            } else {
                var ajax_args = {
                    url: ajaxurl,
                    type: "POST",
                    data: ({step_number : stepNum}),
                    dataType: "text",
                    beforeSend: function(){
                        $this.loader.show();
                    },
                    error: function(){
                        $this.loader.hide();
                    },
                    success: function(res){
                        $this.loader.hide();
                        if(res && res.length>0){
                            selStep.data('hasContent',true);
                            _step($this, selStep).html(res);
                            _showStep($this, stepIdx);
                        }
                    }
                };
                if (ajaxurl_data) {
                    ajax_args = $.extend(ajax_args, ajaxurl_data(stepNum));
                }
                $.ajax(ajax_args);
            }
        }else{
            _showStep($this,stepIdx);
        }
    };

    var _showStep = function($this, stepIdx) {
        var selStep = $this.steps.eq(stepIdx);
        var curStep = $this.steps.eq($this.curStepIdx);
        if(stepIdx != $this.curStepIdx){
            if($.isFunction($this.options.onLeaveStep)) {
                var context = { fromStep: $this.curStepIdx+1, toStep: stepIdx+1 };
                if (! $this.options.onLeaveStep.call($this,$(curStep), context)){
                    return false;
                }
            }
        }
        $this.elmStepContainer.height(_step($this, selStep));
        var prevCurStepIdx = $this.curStepIdx;
        $this.curStepIdx =  stepIdx;
        if ($this.options.transitionEffect == 'slide'){
            _step($this, curStep).slideUp("fast",function(e){
                _step($this, selStep).slideDown("fast");
                _setupStep($this,curStep,selStep);
            });
        } else if ($this.options.transitionEffect == 'fade'){
            _step($this, curStep).fadeOut("fast",function(e){
                _step($this, selStep).fadeIn("fast");
                _setupStep($this,curStep,selStep);
            });
        } else if ($this.options.transitionEffect == 'slideleft'){
            var nextElmLeft = 0;
            var nextElmLeft1 = null;
            var nextElmLeft = null;
            var curElementLeft = 0;
            if(stepIdx > prevCurStepIdx){
                nextElmLeft1 = $this.contentWidth + 10;
                nextElmLeft2 = 0;
                curElementLeft = 0 - _step($this, curStep).outerWidth();
            } else {
                nextElmLeft1 = 0 - _step($this, selStep).outerWidth() + 20;
                nextElmLeft2 = 0;
                curElementLeft = 10 + _step($this, curStep).outerWidth();
            }
            if (stepIdx == prevCurStepIdx) {
                nextElmLeft1 = $($(selStep, $this.target).attr("href"), $this.target).outerWidth() + 20;
                nextElmLeft2 = 0;
                curElementLeft = 0 - $($(curStep, $this.target).attr("href"), $this.target).outerWidth();
            } else {
                $($(curStep, $this.target).attr("href"), $this.target).animate({left:curElementLeft},"fast",function(e){
                    $($(curStep, $this.target).attr("href"), $this.target).hide();
                });
            }

            _step($this, selStep).css("left",nextElmLeft1).show().animate({left:nextElmLeft2},"fast",function(e){
                _setupStep($this,curStep,selStep);
            });
        } else {
            _step($this, curStep).hide();
            _step($this, selStep).show();
            _setupStep($this,curStep,selStep);
        }
        return true;
    };

    var _setupStep = function($this, curStep, selStep) {
        $(curStep, $this.target).removeClass("selected");
        $(curStep, $this.target).addClass("done");

        $(selStep, $this.target).removeClass("disabled");
        $(selStep, $this.target).removeClass("done");
        $(selStep, $this.target).addClass("selected");

        $(selStep, $this.target).attr("isDone",1);

        _adjustButton($this);

        if($.isFunction($this.options.onShowStep)) {
            var context = { fromStep: parseInt($(curStep).attr('rel')), toStep: parseInt($(selStep).attr('rel')) };
            if(! $this.options.onShowStep.call(this,$(selStep),context)){
                return false;
            }
        }
        if ($this.options.noForwardJumping) {
            // +2 == +1 (for index to step num) +1 (for next step)
            for (var i = $this.curStepIdx + 2; i <= $this.steps.length; i++) {
                $this.disableStep(i);
            }
        }
    };

    var _adjustButton = function($this) {
        if (! $this.options.cycleSteps){
            if (0 >= $this.curStepIdx) {
                $($this.buttons.previous).addClass("buttonDisabled");
				if ($this.options.hideButtonsOnDisabled) {
                    $($this.buttons.previous).hide();
                }
            }else{
                $($this.buttons.previous).removeClass("buttonDisabled");
                if ($this.options.hideButtonsOnDisabled) {
                    $($this.buttons.previous).show();
                }
            }
            if (($this.steps.length-1) <= $this.curStepIdx){
                $($this.buttons.next).addClass("buttonDisabled");
                if ($this.options.hideButtonsOnDisabled) {
                    $($this.buttons.next).hide();
                }
            }else{
                $($this.buttons.next).removeClass("buttonDisabled");
                if ($this.options.hideButtonsOnDisabled) {
                    $($this.buttons.next).show();
                }
            }
        }
        // Finish Button
        if (! $this.steps.hasClass('disabled') || $this.options.enableFinishButton){
            $($this.buttons.finish).removeClass("buttonDisabled");
            if ($this.options.hideButtonsOnDisabled) {
                $($this.buttons.finish).show();
            }
        }else{
            $($this.buttons.finish).addClass("buttonDisabled");
            if ($this.options.hideButtonsOnDisabled) {
                $($this.buttons.finish).hide();
            }
        }
    };

    /*
     * Public methods
     */

    SmartWizard.prototype.goForward = function(){
        var nextStepIdx = this.curStepIdx + 1;
        if (nextStepIdx === 1) {
            var prov = $('#prov').val();
            var kab = $('#kab').val();
            var kec = $('#kec').val();

            if (prov === '' || kab === '' || kec === ''){
                if (prov === '') {
                    $('#alert1').remove()
                    $('#alert-1').append(`<div id="alert1"  style="color: red" class="invalid-feedback">
                    Please select your field.
              </div>`)
                } else {
                    $('#alert1').remove()
                }
                if (kab === '') {
                    $('#alert2').remove()
                    $('#alert-2').append(`<div id="alert2" style="color: red" class="invalid-feedback">
                        Please select your field.
                  </div>`)
                } else {
                    $('#alert2').remove()
                }
                if (kec === '') {
                    $('#alert3').remove()
                    $('#alert-3').append(`<div id="alert3" style="color: red" class="invalid-feedback">
                        Please select your field.
                  </div>`)
                } else {
                    $('#alert3').remove()
                }
                return false;
            }
        }
        if (nextStepIdx === 2) {
            var data1 = $('#no_urut').val()
            var data2 = $('#nama_p3a').val()
            var data3 = $('#nama_daerah_irigasi').val()
            var data4 = $('#luas_wilayah').val()
            var data5 = $('#jumlah_p3a').val()
            var data6 = $('#luas_layanan_p3a').val()
            var data7 = $('#keterangan').val()
            if (data1 === '' || data2 === '' || data3 === '' || data4 === '' || data5 === '' || data6 === '' || data7 === '') {
                if (data1 == '') {
                    $('#alert4').remove()
                    $('#alert-4').append(`<div id="alert4" style="color: red" class="invalid-feedback">
                        Please upload your field.
                  </div>`)
                } else {
                    $('#alert4').remove()
                }
                if (data2 == '') {
                    $('#alert5').remove()
                    $('#alert-5').append(`<div id="alert5" style="color: red" class="invalid-feedback">
                        Please upload your field.
                  </div>`)
                } else {
                    $('#alert5').remove()
                }
                if (data3 == '') {
                    $('#alert6').remove()
                    $('#alert-6').append(`<div id="alert6" style="color: red" class="invalid-feedback">
                        Please upload your field.
                  </div>`)
                } else {
                    $('#alert6').remove()
                }
                if (data4 == '') {
                    $('#alert7').remove()
                    $('#alert-7').append(`<div id="alert7" style="color: red" class="invalid-feedback">
                        Please upload your field.
                  </div>`)
                } else {
                    $('#alert7').remove()
                }
                if (data5 == '') {
                    $('#alert8').remove()
                    $('#alert-8').append(`<div id="alert8" style="color: red" class="invalid-feedback">
                        Please upload your field.
                  </div>`)
                } else {
                    $('#alert8').remove()
                }
                if (data6 == '') {
                    $('#alert9').remove()
                    $('#alert-9').append(`<div id="alert9" style="color: red" class="invalid-feedback">
                        Please upload your field.
                  </div>`)
                } else {
                    $('#alert9').remove()
                }
                if (data7 == '') {
                    $('#alert10').remove()
                    $('#alert-10').append(`<div id="alert10" style="color: red" class="invalid-feedback">
                        Please upload your field.
                  </div>`)
                } else {
                    $('#alert10').remove()
                }
                return false;
            }
        }
        if (nextStepIdx === 3) {
            var imgVal = $('#lampiran_tahun_pembentukan').val();
            var imgVal1 = $('#diket_kep_dc').val();
            // var imgVal2 = $('#lampiran_sk_bupati').val();
            // var imgVal3 = $('#lampiran_akte_notaris').val();
            // var imgVal4 = $('#lampiran_pendaftaran').val();
            var data_1 = $('#tahun_pembentukan').val();
            var data_2 = $('#no_sk_bupati').val();
            var data_3 = $('#akte_notaris').val();
            var data_4 = $('#no_pendaftaran').val();
            // imgVal2 === '' || imgVal3 === '' || imgVal4 === '' ||
            if (imgVal === '' || imgVal1 === '' ||  data1 === '' || data2 === '' || data3 === '' || data4 === '') {

                if (data_1 == '') {
                    $('#alert11').remove()
                    $('#alert-11').append(`<div id="alert11" style="color: red" class="invalid-feedback">
                        Please upload your field.
                  </div>`)
                } else {
                    $('#alert11').remove()
                }


                if (imgVal == '') {
                    $('#alert12').remove()
                    $('#alert-12').append(`<div id="alert12" style="color: red" class="invalid-feedback">
                        Please upload your file.
                  </div>`)
                } else {
                    $('#alert12').remove()
                }


                if (imgVal1 == '') {
                    $('#alert13').remove()
                    $('#alert-13').append(`<div id="alert13" style="color: red" class="invalid-feedback">
                        Please upload your file.
                  </div>`)
                } else {
                    $('#alert13').remove()
                }
                if (data_2 == '') {
                    $('#alert14').remove()
                    $('#alert-14').append(`<div id="alert14" style="color: red" class="invalid-feedback">
                        Please upload your field.
                  </div>`)
                } else {
                    $('#alert14').remove()
                }
                // if (imgVal2 == '') {
                //     $('#alert15').remove()
                //     $('#alert-15').append(`<div id="alert15" style="color: red" class="invalid-feedback">
                //         Please upload your file.
                //   </div>`)
                // } else {
                //     $('#alert15').remove()
                // }
                if (data_3 == '') {
                    $('#alert16').remove()
                    $('#alert-16').append(`<div id="alert16" style="color: red" class="invalid-feedback">
                        Please upload your field.
                  </div>`)
                } else {
                    $('#alert16').remove()
                }
                // if (imgVal3 == '') {
                //     $('#alert17').remove()
                //     $('#alert-17').append(`<div id="alert17" style="color: red" class="invalid-feedback">
                //         Please upload your file.
                //   </div>`)
                // } else {
                //     $('#alert17').remove()
                // }
                if (data_4 == '') {
                    $('#alert18').remove()
                    $('#alert-18').append(`<div id="alert18" style="color: red" class="invalid-feedback">
                        Please upload your field.
                  </div>`)
                } else {
                    $('#alert18').remove()
                }
                // if (imgVal4 == '') {
                //     $('#alert19').remove()
                //     $('#alert-19').append(`<div id="aler19" style="color: red" class="invalid-feedback">
                //         Please upload your file.
                //   </div>`)
                // } else {
                //     $('#alert19').remove()
                // }
                return false;
            }

        }
        if (nextStepIdx === 4) {
            var data08 = $('#jumlah_anggota').val();
            var data09 = $('#no_ad_art').val();
            var data10 = $('#lampiran_ad_art').val();
            var data11 = $('#sekretariat').val();
            // var data12 = $('#lampiran_sekretariat').val();
            var data13 = $('#persentase_perempuan_p3a').val();
            var data14 = $('#areal_tersier').val();
            var data15 = $('#pengisian_buku').val();
            var data16 = $('#iuran').val();
            // data12 === '' ||
            if (data08 === '' || data09 === '' || data10 === '' || data11 === '' || data13 === '' || data14 === '' || data15 === '' || data16 === '') {
                if (data08 == '') {
                    $('#alert27').remove()
                    $('#alert-27').append(`<div id="alert27" style="color: red" class="invalid-feedback">
                        Please upload your field.
                  </div>`)
                } else {
                    $('#alert27').remove()
                }
                if (data09 == '') {
                    $('#alert28').remove()
                    $('#alert-28').append(`<div id="alert28" style="color: red" class="invalid-feedback">
                        Please upload your field.
                  </div>`)
                } else {
                    $('#alert28').remove()
                }
                if (data10 == '') {
                    $('#alert29').remove()
                    $('#alert-29').append(`<div id="alert29" style="color: red" class="invalid-feedback">
                        Please upload your field.
                  </div>`)
                } else {
                    $('#alert29').remove()
                }
                if (data11 == '') {
                    $('#alert30').remove()
                    $('#alert-30').append(`<div id="alert30" style="color: red" class="invalid-feedback">
                        Please upload your field.
                  </div>`)
                } else {
                    $('#alert30').remove()
                }
                // if (data12 == '') {
                //     $('#alert31').remove()
                //     $('#alert-31').append(`<div id="alert31" style="color: red" class="invalid-feedback">
                //         Please upload your field.
                //   </div>`)
                // } else {
                //     $('#alert31').remove()
                // }
                if (data13 == '') {
                    $('#alert32').remove()
                    $('#alert-32').append(`<div id="alert32" style="color: red" class="invalid-feedback">
                        Please upload your field.
                  </div>`)
                } else {
                    $('#alert32').remove()
                }
                if (data14 == '') {
                    $('#alert33').remove()
                    $('#alert-33').append(`<div id="alert33" style="color: red" class="invalid-feedback">
                        Please upload your field.
                  </div>`)
                } else {
                    $('#alert33').remove()
                }
                if (data15 == '') {
                    $('#alert34').remove()
                    $('#alert-34').append(`<div id="alert34" style="color: red" class="invalid-feedback">
                        Please upload your field.
                  </div>`)
                } else {
                    $('#alert34').remove()
                }
                if (data16 == '') {
                    $('#alert35').remove()
                    $('#alert-35').append(`<div id="alert35" style="color: red" class="invalid-feedback">
                        Please upload your field.
                  </div>`)
                } else {
                    $('#alert35').remove()
                }
                return false;
            }

        }
        if (nextStepIdx === 5) {
            var t_data1 = $('#operasi').val();
            var t_data2 = $('#partisipatif').val();
            if (t_data1 === '' || t_data2 === '' ){
                if (t_data1 == '') {
                    $('#alert36').remove()
                    $('#alert-36').append(`<div id="alert36" style="color: red" class="invalid-feedback">
                        Please upload your field.
                  </div>`)
                } else {
                    $('#alert36').remove()
                }
                if (t_data2 == '') {
                    $('#alert37').remove()
                    $('#alert-37').append(`<div id="alert37" style="color: red" class="invalid-feedback">
                        Please upload your field.
                  </div>`)
                } else {
                    $('#alert37').remove()
                }
                return false;
            }

        }
        if (nextStepIdx === 6) {
            var p_data1 = $('#pola_tanam').val();
            var p_data2 = $('#usaha_tani').val();
            if (p_data1 === '' || p_data2 === '' ){
                if (p_data1 == '') {
                    $('#alert38').remove()
                    $('#alert-38').append(`<div id="alert38" style="color: red" class="invalid-feedback">
                        Please upload your field.
                  </div>`)
                } else {
                    $('#alert38').remove()
                }
                if (p_data2 == '') {
                    $('#alert39').remove()
                    $('#alert-39').append(`<div id="alert39" style="color: red" class="invalid-feedback">
                        Please upload your field.
                  </div>`)
                } else {
                    $('#alert39').remove()
                }
                return false;
            }
        }
        if (this.steps.length <= nextStepIdx){
            if (! this.options.cycleSteps){
                return false;
            }
            nextStepIdx = 0;
        }
        _loadContent(this, nextStepIdx);
    };

    SmartWizard.prototype.goBackward = function(){
        var nextStepIdx = this.curStepIdx-1;
        if (0 > nextStepIdx){
            if (! this.options.cycleSteps){
                return false;
            }
            nextStepIdx = this.steps.length - 1;
        }
        _loadContent(this, nextStepIdx);
    };

    SmartWizard.prototype.goToStep = function(stepNum){
        var stepIdx = stepNum - 1;
        if (stepIdx >= 0 && stepIdx < this.steps.length) {
            _loadContent(this, stepIdx);
        }
    };
    SmartWizard.prototype.enableStep = function(stepNum) {
        var stepIdx = stepNum - 1;
        if (stepIdx == this.curStepIdx || stepIdx < 0 || stepIdx >= this.steps.length) {
            return false;
        }
        var step = this.steps.eq(stepIdx);
        $(step, this.target).attr("isDone",1);
        $(step, this.target).removeClass("disabled").removeClass("selected").addClass("done");
    }
    SmartWizard.prototype.disableStep = function(stepNum) {
        var stepIdx = stepNum - 1;
        if (stepIdx == this.curStepIdx || stepIdx < 0 || stepIdx >= this.steps.length) {
            return false;
        }
        var step = this.steps.eq(stepIdx);
        $(step, this.target).attr("isDone",0);
        $(step, this.target).removeClass("done").removeClass("selected").addClass("disabled");
    }
    SmartWizard.prototype.currentStep = function() {
        return this.curStepIdx + 1;
    }

    SmartWizard.prototype.showMessage = function (msg) {
        $('.content', this.msgBox).html(msg);
        this.msgBox.show();
    }
    SmartWizard.prototype.hideMessage = function () {
        this.msgBox.fadeOut("normal");
    }
    SmartWizard.prototype.showError = function(stepnum) {
        this.setError(stepnum, true);
    }
    SmartWizard.prototype.hideError = function(stepnum) {
        this.setError(stepnum, false);
    }
    SmartWizard.prototype.setError = function(stepnum,iserror) {
        if (typeof stepnum == "object") {
            iserror = stepnum.iserror;
            stepnum = stepnum.stepnum;
        }

        if (iserror){
            $(this.steps.eq(stepnum-1), this.target).addClass('error')
        }else{
            $(this.steps.eq(stepnum-1), this.target).removeClass("error");
        }
    }

    SmartWizard.prototype.fixHeight = function(){
        var height = 0;

        var selStep = this.steps.eq(this.curStepIdx);
        var stepContainer = _step(this, selStep);
        stepContainer.children().each(function() {
            height += $(this).outerHeight();
        });

        // These values (5 and 20) are experimentally chosen.
        stepContainer.height(height + 5);
        this.elmStepContainer.height(height + 20);
    }

    _init(this);
};



(function($){

$.fn.smartWizard = function(method) {
    var args = arguments;
    var rv = undefined;
    var allObjs = this.each(function() {
        var wiz = $(this).data('smartWizard');
        if (typeof method == 'object' || ! method || ! wiz) {
            var options = $.extend({}, $.fn.smartWizard.defaults, method || {});
            if (! wiz) {
                wiz = new SmartWizard($(this), options);
                $(this).data('smartWizard', wiz);
            }
        } else {
            if (typeof SmartWizard.prototype[method] == "function") {
                rv = SmartWizard.prototype[method].apply(wiz, Array.prototype.slice.call(args, 1));
                return rv;
            } else {
                $.error('Method ' + method + ' does not exist on jQuery.smartWizard');
            }
        }
    });
    if (rv === undefined) {
        return allObjs;
    } else {
        return rv;
    }
};

// Default Properties and Events
$.fn.smartWizard.defaults = {
    selected: 0,  // Selected Step, 0 = first step
    keyNavigation: true, // Enable/Disable key navigation(left and right keys are used if enabled)
    enableAllSteps: false,
    transitionEffect: 'fade', // Effect on navigation, none/fade/slide/slideleft
    contentURL:null, // content url, Enables Ajax content loading
    contentCache:true, // cache step contents, if false content is fetched always from ajax url
    cycleSteps: false, // cycle step navigation
    enableFinishButton: false, // make finish button enabled always
	hideButtonsOnDisabled: false, // when the previous/next/finish buttons are disabled, hide them instead?
    errorSteps:[],    // Array Steps with errors
    labelNext:'Next',
    labelPrevious:'Previous',
    labelFinish:'Finish',
    noForwardJumping: false,
    onLeaveStep: null, // triggers when leaving a step
    onShowStep: null,  // triggers when showing a step
    onFinish: null  // triggers when Finish button is clicked
};

})(jQuery);
