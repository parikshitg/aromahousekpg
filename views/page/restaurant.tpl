{{ template "layout/layout.tpl" . }}

{{ define "title" }} Restaurant - {{ end }}

{{ define "contents" }}

	<div class="container">
		<div class="md-body md-room-detail clearfix">
			<div class="grid_9 md-main">
				<section class="heading-absolute box-sidebar align-center">
					<header class="box-heading heading-large">
						F
						<div class="decription-override">
							<h2 class="h1">Food &amp; Drinks</h2>
						</div>
					</header>
				</section>

				<div class="row">
					<section class="post-body">
						<div class="media">
							<header class="post-header">
								<h2 class="h2">Multi Cuisine Restaurant & Banquet Hall</h2>
							</header>
							{{ with .Restaurant }}
							<figure>
								<img src="{{ image .Image.URI 1900 647 }}" alt="" class="media-object">
							</figure>			
							<div class="media-body">
								<p>{{ htmlString .Richtext }}</p>
							</div>
							{{ end }}
						</div>
					</section>
				</div>

				<div class="row">
					<br><br><br>
				</div>


				<div class="row">
					{{ range $cuisine, $submenu := .FoodMenu }}
					<article class="box box-bullet-list">
						<header class="grid_12 box-heading">
							<h3 class="head headline">{{$cuisine}}</h3>
						</header>

						<div class="box-body">
							{{ range $food := $submenu}}
							<div class="grid_3">
								<div class="col align-center">
									<figure>
										<img src="{{ image $food.Image.URI 320 240 }}">
									</figure>
									<br>
									<header class="box-heading">
										<h5 class="h5">{{$food.Title}}</h5>
										<span class="number number-medium">₹{{$food.Price}}</span>
									</header>
								</div>
							</div>
							{{ end }}
						</div>
					</article>
					{{ end }}
				</div><!-- /.row -->

				<div class="row">
					{{ range $category, $submenu := .DrinkMenu }}
					<article class="box box-bullet-list">
						<header class="grid_12 box-heading">
							<h3 class="head headline">{{$category}}</h3>
						</header>

						<div class="box-body">
							{{ range $i, $drink := $submenu}}
							<div class="grid_3">
								<div class="col align-center">
									<figure>
										<img src="{{ image $drink.Image.URI 320 240 }}" alt="" class="">
									</figure>
									<br>
									<header class="box-heading">
										<h5 class="h5">{{$drink.Title}}</h5>
									<span class="number number-medium">₹{{$drink.Price}}</span>
									</header>									
								</div>
							</div>
							{{ end }}
						</div>
					</article>
					{{ end }}
				</div><!-- /.row -->
			</div><!-- /.md-main -->

			<aside class="grid_3 md-sidebar md-sidebar-pt">
				<section class="box-sidebar">
					<h2 class="h3 header-sidebar"> Menu</h2>
					<ul class="list list-triangle">
						<li><a href="/restaurant">All Items</a></li>
						{{ range $k, $v := .Categories }}
						<li><a href="/restaurant/{{$k}}">{{$v}}</a></li>
						{{ end }}
					</ul>
				</section>
			</aside><!-- /.md-sidebar -->
		</div><!-- /.md-testimonail -->
	</div><!-- /.md-wrapper  -->
{{ end }}
