{{define "header"}}
<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 Transitional//EN" "http://www.w3.org/TR/xhtml1/DTD/xhtml1-transitional.dtd">
<html xmlns="http://www.w3.org/1999/xhtml"><head>
<meta http-equiv="Content-Type" content="text/html; charset=UTF-8">
<meta name="keywords" content="收藏夹,基于新浪APP的应用,网站导航,网址大全"/>
<meta name="description" content="收藏夹,基于新浪APP的应用,网站导航,网址大全"/>
<meta name="author" content="Wasabi" />
<title>网上收藏夹</title>

<!--
<script src="http://ajax.googleapis.com/ajax/libs/jquery/1.3.2/jquery.min.js" type="text/javascript"></script>
-->
<script type="text/javascript" src="js/jquery/jquery.js"></script>

<script type="text/javascript" src="js/jquery/ui/jquery-ui-personalized-1.5.2.min.js"></script>
<!-- jQuery边框-->
<link rel="stylesheet" href="js/jquery/ui/themes/flora/flora.all.css" type="text/css" media="screen" title="Flora (Default)" />

<script src="js/index.js" language="javascript"></script>

<link href="styles/dedestyle.css" rel="stylesheet" type="text/css" />

<meta property="qc:admins" content="21016707131631611006375" />

<script src="lib/thickbox/thickbox.js" type="text/javascript"></script>
<link href="lib/thickbox/thickbox.css" rel="stylesheet" type="text/css" />


<script>
var is_inv		= '0';	//是否邀请
var loginUrl 	= '{$loginUrl}';
</script>


<script>
 $(document).ready(function(){
	  $("#tabsEx1 > ul").tabs({ selected:0});
	  $("#tabsEx2 > ul").tabs({ selected:0});

//2011-07-20
if(is_inv=='1'){
	tb_show("请使用微博账号登录",loginUrl+'?is_tb=1&placeValuesBeforeTB_=savedValues&TB_iframe=true&height=400&width=600');
}

//editItemDiv


//發微博功能
$('#update_weibo').click(function (){

   $.post("index.php?act=update_weibo", {text: $('#text').val()}, function(data){
        alert("" + data);
        $('#text').val('');
  });
});

//給我發評論  對wei收藏提供建议
$('#comment_but').click(function (){
   $.post("index.php?act=comment", {text: $('#messages').val()}, function(data){
        alert("" + data);
        $('#messages').val('');
  });
});


//检查分类select
$('#item_select').change( function () {

if(this.value=='add'){

   var name = prompt('请输入新分类的名字。');

   if(name!=null){
	var data =  {title: name};

	$.post("/item/edititem",data, function(data){
		//alert("Data Loaded: " + data);

		 $("<option value='"+ data + "'>"+name+"</option>").appendTo("#item_select")//添加下拉框的option

		 $('#item_select').val(data);
  });

   }
   else
		$('#item_select').attr('selectedIndex',0);
}

  //如果分类Select 的值大于0，则得到相应的分类ID，那得到相应分类的最后一个item的seq
  else if(this.value > 0){
        settingLastSeq(this.value);
  }


});



$('#addItemButton').click( function () {

	var data =  { title: $('#title').val(), description: $('#item_description').val(),sort:$('#sort').val()};

	$.post("index.php?act=add_item",data, function(data){
		alert("Data Loaded: " + data);
	});

 });


// 编辑收藏
 $('.hyplink_td').dblclick( function () {

	hyplink = $(this).find('.hyplink');	//得到相關的連接

	$('#link_id').val(hyplink.attr('id'));
	$('#name').val(hyplink.text());
	$('#href').val(hyplink.attr('href'));
	$('#description').val(hyplink.attr('title'));

	$('#seq').val(hyplink.attr('seq'));


     $('#item_select').val(hyplink.attr('item_id'));
//     $('#item_select_div').hide(); //把分类隐藏

//	$(this).find('.edit_link').click();

      tb_show("编辑收藏","#TB_inline?&width=380&height=130&amp;inlineId=linkContent");
});

// 编辑收藏
 $('.edit_item').dblclick( function () {

	var id		= $(this).attr('id');
	var name	= $(this).text();
	var sort		= $(this).attr('sort');

	/*********************************************************/

	$('#item_id').val(id);

	$('#it_title').val(name);

	$('#it_seq').val(sort);

	$('#it_page').val($(this).attr('page'));

      tb_show("编辑收藏","#TB_inline?&width=380&height=130&amp;inlineId=editItemDiv");
});



   //根据相应的分类ID，那得到相应分类的最后一个item的seq
   function settingLastSeq (cat_id){
     //var cat_id = this.value;

        var seq = $('#table_'+ cat_id + ' tr:last td:last a').attr('seq');

        $('#seq').val(Number(seq)+1);

   }

function clean_hyplink_form(){
    $('#link_id').val('');
	$('#name').val('');
	$('#href').val('');
	$('#description').val('');
	$('#seq').val('');
}


//新增收藏
$('#add_link_but').click( function(){
   clean_hyplink_form();

  //相应的分类ID，那得到相应分类的最后一个item的seq
   settingLastSeq ( $('#item_select').val());

 tb_show("新增收藏","#TB_inline?&width=380&height=180&inlineId=linkContent");

   return false;

});


$('#edit_link_but').click( edit_link);

//新增和修改连接
function edit_link(){

var data =  { link_id: $('#link_id').val(), name: $('#name').val(),href:$('#href').val(),description:$('#description').val(),seq:$('#seq').val(),item_id:$('#item_select').val()};

//如果link_id为空，即是新增收藏
if(data.link_id==''){
//查看该分类的最后一行有多少个td
var td_length = $('#table_'+data.item_id).find('tr.tr_links:last-child td').length;

  var str='<td class="hyplink_td"><a title="'+data.description+'" seq="10" item_id="15" id="110" target="_blank" href="'+data.href+'" class="hyplink">'+ data.name +'</a></td>';

if(td_length==4){
  str = '<tr class="tr_links">'+str+'</tr>'

  $('#table_'+data.item_id).append(str);
}
else
  $('#table_'+data.item_id).find('tr.tr_links:last-child').append(str);
}

  var link = $('#'+data.link_id);
  link.attr('href', data.href);
  link.attr('title', data.description);
  link.attr('seq', data.seq);
  link.text(data.name);

  //分類 item_select

 tb_remove();

$.post("/item/editlink",data, function(data){
   alert(data);
  });
}


$('#edit_item_but').click( edit_item);

//新增和修改连接
function edit_item(){

var data =  { item_id: $('#item_id').val(), title: $('#it_title').val(), sort:$('#it_seq').val() , page:$('#it_page').val() };

  //分類 item_select

 tb_remove();

$.post("/item/edititem",data, function(data){
   // alert("Data Loaded: " + data);
  });
}



/*********************************************************************************/

});

function showLogin_tb(){

	//if($("#share").attr('checked')==undefined)  alert('fdfdsf');

tb_show("请使用微博账号登录",loginUrl+'?is_tb=1&placeValuesBeforeTB_=savedValues&TB_iframe=true&height=400&width=600');
}

/**刪除連接*/
function delete_link(){
    var data =  { link_id: $('#link_id').val()};

    if(data.link_id=='') return false;

     $('#'+data.link_id).parent().html('');

    $.post("/item/deletelink",data, function(data){
    //alert("Data Loaded: " + data);
        tb_remove();
      });


}
</script>
</head>
<body>

<div id="page">
<div class="top cbody">

<div class="toplogo">

<a href="{$HTTP_HOME}"><IMG alt="Wasabi" src="http://static.sae.sina.com.cn/image/logo.beta.png" height="55" width="185"></a>

</div>

<div class="topbanner">
	<img src="indexbb/ysl1.gif" width="468" height="60" />
</div>

<div class="toplink">

{{if $.user_id}}
<div class="user_head">
  <div id="pop_1" class="picborder_r lf">
  <a href="#">
     <img alt="" pop="true" src="{{$.user_info.AvatarsPath}}" class="person_icon"></a></div>
   <div class="lf">
	<p class="font_14"><b>{{$.user_info.Nickname}} </b></p>

	<p><em><a href="/login/logout">退 出</a> </em></p></div>
  </div>
{{else}}
<div style="margin-top:15px;">
  <a title="点击进入授权页面" alt="点击进入授权页面" href="{$loginUrl}"><img src="images/weibo_connect.gif"></a><br />
</div>
{{end}}

</div>
</div><!--top-->


<div class="topmenu cbody">
<ul>
    <li><a href="{$HTTP_HOME}">首页</a> </li>
	<li><a href="http://wasabi.sinaapp.com/" target="_blank">我的博客</a> </li>
	<li><a href="http://wasabi.iteye.com/" target="_blank">ITeys bolg</a> </li>
	<li><a href="http://wawiki.sinaapp.com/" target="_blank">我的wiki</a> </li>
  <li><a href="https://www.google.com/reader/" target="_blank">GoogleReader</a> </li>
  <li><a href="http://reader.youdao.com/" target="_blank">有道閱讀</a></li>

{{if $.is_edit}}
  <li class="spece"></li>
  <li class="funs"><a id="add_link_but" href="#" class="" target="_blank">新增收藏</a> </li>
{{end}}
<li class="funs"><a id="" href="http://wasa.sinaapp.com/?act=inv" class="" target="_blank">邀请好友注册</a> </li>
<li class="funs"><a id="" href="http://wasa.sinaapp.com/guide/#!prettyPhoto[gallery2]/0/" class="" target="_blank">新手指南</a> </li>
</ul>
</div>
{{end}}