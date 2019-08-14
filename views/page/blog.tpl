{{ template "layout/layout.tpl" . }}

{{ define "title" }} Blogs -  {{ end }}

{{ define "contents" }}
	<!-- Body -->
	<div class="md-body-wrapper">
		<div class="container">
			<div class="md-body md-blog clearfix">
				<div class="grid_9 md-main">
					<section class="heading-absolute box-sidebar align-center">
						<header class="box-heading heading-large">
							B
							<div class="decription-override">
								<h2 class="h1">our blog</h2>	
							</div>
						</header>
					</section>

					{{ if .BlogList }}
					<div class="list-article">
						{{ range $post := .BlogList }}
						<article class="media">
							<!-- <h2>tranhieu</h2> -->
							<figure class="pull-left">
								<a href="/blog/{{$post.Slug}}"><img src="{{ image $post.Image.URI 349 194 }}" alt="" class="media-object"></a>
							</figure>
							<div class="media-body">
								<header class="media-header">
									<div class="meta-header">
										<span class="time">{{ month $post.Date }} {{ day $post.Date }}, {{ year $post.Date }} </span>
										<span class="meta-divider">|</span>
										<p class="tags-post">
											Tags:
											{{ range $i, $tag := .Tags -}}
											<a href="#" class="tag-post">{{ if gt $i 0 }}, {{ end }}{{ $tag }}</a>
											{{- end }}
										</p>
									</div>
									<h3 class="h4"><a href="/blog/{{$post.Slug}}">{{ $post.Title }}</a></h3>
								</header>
								<p class="media-content">{{ htmlString $post.Body }}</p>
								<p class="meta-post">
									<span class="meta-pers"><i class="icon icon-view"></i>{{ $post.Author }}</span>
								</p>
							</div>
						</article>
						{{ end }}
					</div>
					<ul class="pagination">
						<li class="first"><a href="#">First</a></li>
						<li><a href="#">1</a></li>
						<li><a href="#">2</a></li>
						<li><a href="#">3</a></li>
						<li class="last"><a href="#">Last</a></li>
					</ul>
					{{ end }}

					{{ with .Post }}
						<div class="row">
							<section class="post-body">
								<div class="media">
									<header class="post-header">
										<span class="time">{{ month .Date }} {{ day .Date }}, {{ year .Date }} </span>
										<h2 class="h2">{{.Title}}</h2>
									</header>
									<figure>
										<img src="{{ image .Image.URI 881 352 }}" alt="" class="media-object">
									</figure>
									<div class="media-body">
										<p>{{ htmlString .Body}}</p>
									</div>
								</div>
							</section>
							<footer class="post-footer clearfix">
								<div class="footer-left">
									<p class="author-post">
										Author: <span>{{.Author}}</span>
									</p>
									<p class="tags-post">
										Tags:
										{{ range $i, $tag := .Tags -}}
										<a href="#" class="tag-post">{{ if gt $i 0 }}, {{ end }}{{ $tag }}</a>
										{{- end }}
									</p>
								</div>
								<div class="footer-right">
									Share this:
									<a href="#"><i class="icon icon-facebook"></i></a>
									<a href="#"><i class="icon icon-twitter"></i></a>
									<a href="#"><i class="icon icon-google"></i></a>
									<a href="#"><i class="icon icon-dribbble"></i></a>
								</div>
							</footer>
						</div>					
					{{ end }}

				</div><!-- /.md-main -->
				<aside class="grid_3 md-sidebar md-sidebar-pt">
					<section class="box-sidebar">
						<h3 class="h3 header-sidebar">Categories</h3>
						<ul class="list-cate list list-triangle list-uppercase">
							<li><a href="#">Cras condimentum</a></li>
							<li><a href="#">Nibh et pellentesque</a></li>
							<li><a href="#">Leo nibh gravida velit</a></li>
							<li><a href="#">Tortor augue a eros</a></li>
							<li><a href="#">libero</a></li>
						</ul>
					</section>

					<section>
						{{ template "partial/check-sidebar.tpl" . }}
					</section>

					<section class="box-sidebar">
						<h3 class="h3 header-sidebar">Tags</h3>
						<ul class="list-tags clearfix">
							<li><a href="#">Suspendisse</a></li>
							<li><a href="#">justo</a></li>
							<li class="active-tag"><a href="#">consectetur</a></li>
							<li><a href="#">tristique tempus </a></li>
							<li><a href="#">laoreet in</a></li>
							<li><a href="#">pharetra</a></li>
						</ul>
					</section>
				</aside><!-- /.md-sidebar -->
			</div><!-- /.md-testimonail -->
		</div>
	</div>

{{ end }}