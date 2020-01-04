<!DOCTYPE>
<html>
    <head>
        <meta charset="utf-8"/>
        <title>2k745</title>
		<link rel="stylesheet" href="https://www.layuicdn.com/layui/css/layui.css" media="all">
        <style>
            h1{text-align:center;}
            .big-pic{border:1px solid black;text-align:center;width:100%;} 
			.big-pic ul {padding:0 10px 5px;text-align:left;overflow:auto;display:block;height:90%}
			.big-pic ul li {list-style:none;list-style-type:none;margin:0px;}
            .big-pic .pic-frame{border:1px solid red;display:inline-block;padding:5px;}
			.big-pic .pic-frame-left{display:inline-block;float:left;width:28%;height:260;}
			.big-pic .pic-frame-right{padding:30px 5px;display:inline-block;float:right;width:28%;height:200;}
            .big-pic .pic-frame-right .send-content {height:70%;display:block;width:100%;margin-bottom:10px;}
			.big-pic .pic-frame-right button {width:150px;float:right;margin-right:10px;}
			
			.main ul {text-align:center;}
            .main ul li{width:100px;border:1px red solid;list-style:none;margin:10px;display:inline-block;padding:2px;margin-right:100px;}
            .main ul li img{width:100px;height:100px}
            .main ul li button{margin:8px;}
            .main ul li .user-name{color:green;}
           
        </style>
        <script src="/static/js/jquery.mini.js"></script>
    </head>
    <body>
        <div id="page">
            <div class="top">
                <h1>2k745投票系统</h1>
                <div class="big-pic">
                		<div class="pic-frame-left">
						<div class="message">
							<ul>
								<li id="0"></li>
							</ul>
						</div>	
					</div>
					<div class="pic-frame"><img src="{{.QiNiuPath}}/{{.MaxVote.Img}}" width="300" height="250"/><div class="user-vote">{{.MaxVote.Count}}</div></div>
					<div class="pic-frame-right">
						<textarea class="send-content" placeholder="发送内容"></textarea>
						<button id="send">发送</button>
					</div>
                </div>
            </div>
            <div class="main">
                <ul>
                    {{range $index, $item := .Votes}}
                        <li id="{{$item.Id}}" class="user-node" ><img src="{{$.QiNiuPath}}/{{$item.Img}}"/><div class="user-name">{{$item.Name}}</div><div class="user-vote">{{$item.Count}}</div><button class="user-click">投票</button></li>
                    {{end}}
                    <!--<li><img src="/static/img/top_pic.jpg"/><div class="user-name">姓名</div><button>投票</button></li>
                    <li><img src="/static/img/top_pic.jpg"/><div class="user-name">姓名</div><button>投票</button></li> -->
                </ul>
            </div>
            <div class="footer"></div>
        </div>
    </body>
	<script src="https://www.layuicdn.com/layui/layui.js"></script>
    <script>
        $('.user-click').click(function(){
            var id = $(this).parent('li').attr('id') 
			//投票
            //$.get('/add', {id:id});
			$.ajax({
			  url: "/add",
			  async: false,
			  data: "id="+id,
			 });
            location.href='/search?' + Math.random();
        })
		
		//获取聊天消息
		/*function getMessageList() {
			var lastId = $(".message ul li:last").attr('id');
			$.ajax({
			  url: "/getLastMessage",
			  async: false,
			  data: "lastId="+lastId,
			  success: function(msg){
     			//var megs = eval('(' + msg + ')');
				for(i in msg) {
					//console.log(msg[i])
					$(".message ul").append("<li id=" + msg[i].Id + ">"+ msg[i].Ip + ": " + msg[i].Content + "</li>");
				}
   			  }	
			 });
			var div = $(".message ul");
			div[0].scrollTop = div[0].scrollHeight;
			
		}
		getMessageList();
		setInterval("getMessageList()","400");
		*/
		var sock = null;
        var wsuri = "ws://127.0.0.1:10001/getMessage";
        sock = new WebSocket(wsuri);
        sock.onmessage = function(e) {
			var msg = eval('(' + e.data + ')');
			for(i in msg) {
				//console.log(msg[i])
				$(".message ul").append("<li id=" + msg[i].Id + ">"+ msg[i].Ip + ": " + msg[i].Content + "</li>");
			}
			var div = $(".message ul");
			div[0].scrollTop = div[0].scrollHeight;
        }
		
		setInterval(function(){
			var lastId = $(".message ul li:last").attr('id');
			sock.send(lastId);
		},"800");
		
 
		//发送消息
		$("#send").click(function(){
            var content = $(".send-content").val();
			$.ajax({
			  type:"post",
			  url: "/addMessage",
			  async: false,
			  data: "content="+content,
			  success: function(msg){
				 layui.use(['layer'], function() {
      			 	var layer = layui.layer;
      				layer.msg('发送成功');
    		    });
			  }
			 });
        })
    </script>
</html>
