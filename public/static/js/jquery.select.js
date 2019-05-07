(function($){ 
	$.fn.s_select = function(options){ 
		var defaults = { 
			inputID: 's_input_'+$(this).attr('id'),
			inputName: 'hospitalCodeName',
			selectID: $(this).attr('id')
		};
		var opts = $.extend(defaults, options); 
		var $_this = $(this);
//		var inputValue = $(this + ":selected").html();
		var inputValue = $("#"+opts.selectID + " :selected").html();
		var inputValue2 = $("#"+opts.selectID + " :selected").val();
		
		if(!inputValue) inputValue = "--未选择--";
		var selHtml = $(' <input id="' + opts.inputID + '" type="text" name="' + opts.inputName + '" value="' + inputValue + '">' 
						+ ' <input type="hidden" id="na_' + opts.selectID + '" type="text" name="na_' + opts.selectID + '" value="' + inputValue2 + '">');
		$(selHtml).insertBefore($_this);
		$('#' + opts.inputID).focus(function(){
			showSelect();
			// showDivStation($(this)[0], true,'hospitalCode');
		});
		$('#'+opts.inputID).keyup(function(e){
			showSelect();
			similarFind($(this)[0]);
			e ? e : window.event;
			if(e.keyCode == 13){
				e.preventDefault();
				selectStation($_this[0]);
			}
		});
		$('#'+opts.inputID).click(function(){
			removeNull($(this)[0]);
		});
		$_this.click(function(){
			selectStation($_this[0]);
		});
		$('#' + opts.inputID).dblclick(function(){
			$_this.hide();
		});
		//function codes in here 
		var whichText = $('#' + opts.inputID)[0];  
		var whichValue = $('#na_' + opts.selectID)[0];
		function selectStation(obj){  
			var objSelStation = obj;  
			if (obj.selectedIndex != -1) {  
				var stationName = obj.options[obj.selectedIndex].text;  
				whichText.value = stationName;
				whichValue.value = obj.options[obj.selectedIndex].value ;
				$_this[0].value=stationName;
				$_this.css({display:'none'});
			}  
			var stobj = document.getElementById(opts.selectID);  
		}  
		function showSelect(){//显示下拉框
			$_this.css({display: 'block'});
		}
		var pageD = 0, pageU;   
		function similarFind(txtObj) {  //查找类似文本，参数为输入节点及选择后赋值的节点名
			var curStationName = txtObj.value;  //接收到输入的值
			var objSelStation = document.getElementById(opts.selectID);  //找到选择节点
			var stationLength = objSelStation.options.length;  //选择选项的数量
			var flag = true;  
			pageU = pageD;  
			//从起始的文字匹配 用text中的数据跟下拉框中的数据  
		    for(var i = 0; i < stationLength; i++){  
		        var stationName = objSelStation.options[i].text;  //循环选项的值，看是否和输入的值有类似
		        var re = new RegExp("^" + curStationName);  
		        if (stationName.match(re)) {  
		            if (i<stationLength - 10) {  
		                objSelStation.selectedIndex = i + 10;  
		            }  
		            objSelStation.selectedIndex = i;  
		            pageD = i;  
		            pageU = i;  
		            flag = false;  
		            break;  
		        }  
		    }  
		    ////从文字中匹配 用text中的数据跟下拉框中的数据  
		    if(flag) {  
		        for (var i=0; i<stationLength; i++) {  
		            var stationName = objSelStation.options[i].text;  
		            var re2 = new RegExp("^.*" + curStationName+'.*$');  
		            if (stationName.match(re2)) {  
		                if (i<stationLength - 10) {  
		                    objSelStation.selectedIndex = i + 10;  
		                }  
		                objSelStation.selectedIndex = i;  
		                pageD = i;  
		                pageU = i;  
		                break;  
		            }  
		        }  
		    }  
		    //响应下移键  
			var e=event||window.event;
		    if(e.keyCode==40) {  
		        pageD++;  
		        if(pageD==objSelStation.options.length) pageD=0;  
		        txtObj.value=objSelStation.options[pageD].text ;  
		        objSelStation.selectedIndex = pageD;  
		    }  
		    //响应上移键  
		    if(e.keyCode==38) {  
		        --pageU;  
		        if(pageU<0) pageU=objSelStation.options.length-1;  
		        txtObj.value = objSelStation.options[pageU].text;  
		        objSelStation.selectedIndex = pageU;  
		    }  
		}  
				
		function removeNull(node) {  //点击输入框文字'--未选择--'去掉
	        if(node.value=='--未选择--')   
	        node.value = '';  
	    }  
	};
})(jQuery);