function coutStrNum(id,showId,limit_Len){
	
	     var thisInput=document.getElementById(id),count_t_f =true; 
				
	     String.prototype.len=function(){ 
				return this.replace(/[^\x00-\xff]/g,"**").length; 
			} 
			
	     function GetCharLength(str){ //统计字数
			var iLength = 0; 
			for(var i = 0;i<str.length;i++) 
			{ 
				if(str.charCodeAt(i) >255) 
				{ 
				iLength += 2; 
				} 
				else 
				{ 
				iLength += 1; 
				} 
			} 
			return iLength; 
			}

			 function CutStr(Str,Len) //显示剩余字数，同时把多余的截断
			{ 
				var CurStr="";
				for(var i = 0;i<Str.length;i++) 
				{ 
					CurStr += Str.charAt(i); 
					if(CurStr.len()>Len) //如果字数超过限制，不让提交
					{ 
				    document.getElementById(showId).innerHTML=" 剩余："+(Len-CurStr.len())+"字节";
				    count_t_f=false;
					
					} 
					
					document.getElementById(showId).innerHTML=" 剩余："+(Len-CurStr.len())+"字节"; 
					
				} 
				
				
			} 

			 

			
			document.getElementById(id).onkeyup=function(){//绑定onkeyup事件
					var Str=thisInput.value;
			     	CutStr(Str,limit_Len);
			     	
			     }
			
	
			document.getElementById(id).onkeyup();
			if(!count_t_f){
				return false; 
			}else{
				return true ; 
			}
	
	}