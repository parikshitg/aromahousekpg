{{ template "layout/layout.tpl" . }}

{{ define "title" }} About -  {{ end }}

{{ define "contents" }}
	<div class="container">
		<div class="md-body md-standar-post clearfix">
			<div class="grid_12 md-main">
				<div class="row">
					<section class="post-body">
						{{ with .About}}
						<div class="media">
							<header class="post-header">
								<!--<span class="time">March 7, 2013 </span>-->
								<h2 class="h2">Aroma House</h2>
							</header>
							<figure>
								<img src="{{ image .Image.URI 1900 800 }}" alt="" class="media-object">
							</figure>
							<div class="media-body">
								<p>{{ htmlString .Richtext }}</p>
							</div>
						</div>
						{{ end }}
						<br><br><br>

						{{ with .Chef }}
						<div class="media media-col">
							<h3 >Bikas Lama (Owner & Master Chef)</h3>
							<br>
							<figure class="pull-left">
								<img src="{{ image .Image.URI 400 600 }}" alt="" class="media-object">
							</figure>
							<div class="media-body">
								<p>{{ htmlString .Richtext }}</p>
								<br>
								<br>

								<h4 class="h4">Skills & Expertise</h4>
								<div class="grid_3">
									<ul class="list list-triangle list-triangle-color">
										<li></li>
										{{ range $i, $v := .List }}
											{{ if isOdd $i }}
												<li>{{$v}}</li>
											{{ end }}
										{{ end }}
									</ul>
								</div>
								<div class="grid_3">
									<ul class="list list-triangle list-triangle-color">
										<li></li>
										{{ range $i, $v := .List }}
											{{ if isEven $i }}
												<li>{{$v}}</li>
											{{ end }}
										{{ end }}
									</ul>
								</div>										
							</div>
						</div>
						{{ end }}
					</section>
	
				</div>
				
			</div><!-- /.md-main -->
		</div><!-- /.md-testimonail -->
	</div><!-- /.md-wrapper  -->
{{ end }}
