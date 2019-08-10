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
								<img src="http://placeholder.mac/1900x800.png" alt="" class="media-object">
							</figure>
							<div class="media-body">
								{{ htmlString .Richtext }}
							</div>
						</div>
						{{ end }}
						<br><br><br>

						{{ with .Chef }}
						<div class="media media-col">
							<h3 >Bikas Lama (Owner & Master Chef)</h3>
							<br>
							<figure class="pull-left">
								<img src="http://placeholder.mac/400x600.png" alt="" class="media-object">
							</figure>
							<div class="media-body">
								{{ htmlString .Richtext }}
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
								<!--<p>Donec eget tellus eget quam facilisis egestas. Maecenas molestie turpis at tincidunt commodo. Aenean et sapien in lacus sagittis elementum a nec. Nunc at tristique erat. Vestibulum at augue id ante placerat porta ut sit amet mi. Phasellus venenatis odio nec massa dictum, et bibendum justo iaculis. Sed lobortis diam et massa vestibulum, id sagittis purus malesuada. Nullam aliquam feugiat nisi. Aliquam semper, lectus id porttitor dignissim, sem risus sollicitudin turpis, a rhoncus diam arcu non quam. Interdum et malesuada fames ac ante ipsum primis in faucibus. Cras fringilla sollicitudin metus et mattis. Sed blandit egestas ante, nec accumsan lectus sollicitudin ac.</p>
								<p>Donec eget tellus eget quam facilisis egestas. Maecenas molestie turpis at tincidunt commodo. Aenean et sapien in lacus sagittis elementum a nec. Nunc at tristique erat. Vestibulum at augue id ante placerat porta ut sit amet mi. Phasellus venenatis odio nec massa dictum, et bibendum justo iaculis. Sed lobortis diam et massa vestibulum, id sagittis purus malesuada. Nullam aliquam feugiat nisi. Aliquam semper, lectus id porttitor dignissim, sem risus sollicitudin turpis, a rhoncus diam arcu non quam. Interdum et malesuada fames ac ante ipsum primis in faucibus. Cras fringilla sollicitudin metus et mattis. Sed blandit egestas ante, nec accumsan lectus sollicitudin ac.</p> -->										
							</div>
						</div>
						{{ end }}
					</section>
	
				</div>
				
			</div><!-- /.md-main -->
		</div><!-- /.md-testimonail -->
	</div><!-- /.md-wrapper  -->
{{ end }}
