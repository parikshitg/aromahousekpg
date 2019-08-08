{{ template "layout/layout.tpl" . }}

{{ define "title" }} Gallery - {{ end }}

{{ define "contents" }}
	<div class="container">
		<div class="md-body clearfix">
			<div id="md-news-deal" class="md-news-deal masonry-container masonry">
				
				{{ range $i, $pic := .Gallery}}
				{{ if eq $i 0 }}
				<article class="media media-center">
					<figure>
						<a href="room-detail.html"><img src="http://placeholder.mac/800x600.png" alt="" class="media-object" width="380" height="auto"></a>
					</figure>
					<div class="media-body">
						<h3 class="media-header h4"> {{$pic.Title}}</h3>
						<p class="media-content">{{$pic.Text}}</p>
					</div>
				</article>
				{{ end }}
				{{ end }}

				<article class="media media-center">
					<header class="box-heading heading-large">
						G
						<div class="decription-override">
							<h2 class="h1">Gallery</h2>
							<p>what's happening</p>
						</div>
					</header>
					<div class="media-body">
						<p class="media-content">Mauris egestas, tellus sed venenatis tincidunt, odio diam iaculis augue, nec tincidunt enim odio id arcu. Ut pellentesque, quam ut sagittis adipiscing</p>
					</div>
				</article>


				{{ range $i, $pic := .Gallery}}
				{{ if eq $i 1 }}
				<article class="media media-center">
					<figure>
						<a href="room-detail.html"><img src="/images/paragliding2.jpg" alt="" class="media-object" width="380" height="auto"></a>
					</figure>
					<div class="media-body">
						<h3 class="media-header h4"> {{$pic.Title}}</h3>
						<p class="media-content">{{$pic.Text}}</p>
					</div>
				</article>
				{{ end }}
				{{ end }}


				{{ range $i, $pic := .Gallery}}
				{{ if gt $i 1 }}
				<article class="media media-center">
					<figure>
						<a href="room-detail.html"><img src="/images/paragliding2.jpg" alt="" class="media-object" width="380" height="auto"></a>
					</figure>
					<div class="media-body">
						<h3 class="media-header h4"> {{$pic.Title}}</h3>
						<p class="media-content">{{$pic.Text}}</p>
					</div>
				</article>
				{{ end }}
				{{ end }}

			</div>
		</div><!-- /.md-testimonail -->
	</div><!-- /.md-wrapper  -->

{{ end }}

{{ define "js" }}
<script type="text/javascript" src="/js/jquery.masonry.js"></script>
<script type="text/javascript">
		$(document).ready(function(){
			"use strict";
			var newdeal = document.querySelector('#md-news-deal');
			var loadmore = document.querySelector('#load-more');
			
			var  msnry = new Masonry(newdeal,{
				itemSelector: '.media-center',
		        columnWidth: 400,
		        isFitWidth: true
			});

			$("#load-more").bind("click",function(event){
				event.preventDefault();
				$.get("loadmore.tpl",function(response){
					var data = $(response).filter("article.media-center");
					console.log(data);
					$("#md-news-deal").append(data);
					msnry.appended(data);
				},'html')
			})


		 jQuery(function(){
		     jQuery('.view-content').masonry({
		          // options
		         itemSelector : '.views-row'
		      });
		  });

 
		});
	</script>
{{ end }}