{{template "header" .}}
<div id="content">
<div id="tabsEx1" style="width: 620px; height: 730px;" class="table_style">
<ul style="height: 30px;">
	<li><a href="#fragment-1"><span>常用一</span></a></li>
	{{if gt $.item_length 14}}<li><a href="#fragment-2"><span>常用二</span></a></li>{{end}}
	{{if gt $.item_length 28}}<li><a href="#fragment-3"><span>常用三</span></a></li>{{end}}
</ul>

{{range $key,$val := .item_list}}

    {{if eq $key 0}}<div id="fragment-1">
    {{else if eq $key 14}}
    </div><div id="fragment-2">
    {{else if eq $key 28}}
    </div><div id="fragment-3">
    {{end}}

        {{if $val.Link_list}}
    	<table width="100%" id="table_{{$val.Id}}"><!-- 分類名 -->
    		<tr> <th {{if $.is_edit}}class="edit_item" id="{{$val.Id}}" sort="{{$val.Sort}}" page="{{$val.Page}}"{{end}} colspan="4" align="left">{{$val.Title}}</th></tr>

            {{range $link_key, $link := $val.Link_list}}
                {{if eq 0 ($link_key | molding)}} <tr class="tr_links">{{end}}
                    		<td class="{{if $.is_edit}}hyplink_td{{end}}">
                    			<a class="hyplink" href="{{$link.Href}}" target="_blank" id="{{$link.Id}}" item_id="{{$val.Id}}" seq='{{$link.Seq}}' title="{{$link.Description}}" alt="{{$link.Description}}" >{{$link.Name}}</a>
                 {{if eq 1 ($link_key | modend)}}</tr>{{end}}
            {{end}}
        </table><br />
        {{end}}
{{end}}
</div>
	<div><a class="hyplink" href="http://www.beian.miit.gov.cn/" target="_blank" alt="">粤ICP备16031203号</a></div>
</div>

<div id="tabsEx2" style="height: 730px;" class="table_style">
<ul style="height: 30px;">
	<li><a href="#fragment-4"><span>收藏頁</span></a></li>
	{{if $.is_edit}}
		<li><a href="#fragment-5"><span>微博相關</span></a></li>
	{{end}}
	<li><a href="#fragment-6"><span>Three</span></a></li>
</ul>

{{template "fav" .}}
<div id="fragment-6">
</div>
</div>


<div id="linkContent" style="display:none;">
	<form action="{{$.curr_page}}?act=edit_link" method="post">
		<input type="hidden" id="link_id" name="link_id" value="" />
		<input type="hidden" id="item_id" name="item_id" value="" />

		<div style="">
			<div id="item_select_div" class="leftdiv" >
			所属分类：
			<select id="item_select" name="item_select" >
				{{range $key,$item := .item_list}}
			        <option value="{{$item.Id}}"> {{$item.Title}}</option>
			    {{end}}

			<option value="0" class="">------------------</option>
            <option value="add" class="thickbox">新增分类</option>
			</select>

			<span style="margin-left:80px;"><a hef="#" onclick="delete_link();return false;"> 删除该网址</a></span>
                 <br /></div>

			<div class="leftdiv">名称： <input type="text" id="name" name="name" value="" size="25" /></div>
			<div class="leftdiv">网址： <input type="text" id="href" name="href" value="" size="25" /></div>
			<div class="leftdiv">描述： <input type="text" id="description" name="description" value="" size="25" />
			排序： <input type="text" id="seq" name="seq" value="" size="2" />
			</div>

			<div class=" butt" style="text-align:center; margin-top: 15px;">
			<input id="edit_link_but" type="button" value="提 交" />
			&nbsp;&nbsp;&nbsp;&nbsp;
			<input type="button" value="取 消" onclick="tb_remove();" />
			</div>
	        </div></form>
</div>

<div id="editItemDiv" style="display:none;">
	<div style="">
	<div class="leftdiv">名称： <input type="text" id="it_title" name="it_title" value="" size="25" />
	<input type="hidden" id="item_id" name="item_id" value="">

	</div>

	Page： <input type="text" id="it_page" name="it_page" value="" size="2" />

	排序： <input type="text" id="it_seq" name="it_seq" value="" size="2" />
	</div>
	<div class=" butt" style="text-align:center; margin-top: 15px;">
		<input id="edit_item_but" type="button" value="提 交" />&nbsp;&nbsp;&nbsp;&nbsp;
		<input type="button" value="取 消" onclick="tb_remove();" />
	</div>
</div></div>
</div>


</div><!--END contentDIV-->
</div><!--END page -->
<div></div>
</body>
</html>