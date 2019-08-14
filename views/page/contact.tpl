{{ template "layout/layout.tpl" . }}

{{ define "title" }} Contact - {{ end }}

{{ define "css" }} 
<link rel="stylesheet" type="text/css" href="/css/libs/magnific-popup.css">
<style>
		#map_canvas {
	        /*width: 500px;*/
	        height: 820px;
	      }

	</style>
{{ end }}

{{ define "contents" }}
	{{ with .Contact }}
		<section class="md-contact">
			<figure class="bg-contact">
				<img src="{{ image .Image.URI 1560 814 }}" alt="">
				<div class="bg-pattern"></div>
			</figure>
			<div class="contact-wrap">
				<div class="layout-left">
					<header class="box-heading heading-large">
						C
						<div class="decription-override">
							<h2 class="h1">{{ .Title }}</h2>
							<p> {{ .Subtitle }} </p>
						</div>
					</header>
					<p id="contact-content" class="description">{{ .Richtext }}</p>
					<div class="row form-contact">
						<form  class="label-placeholder" action="/contact" method="POST">
							<div class="row clearfix">
								<div class="form-group">
									<input  type="text" name="name" class="input-control" placeholder="Name">
								</div>
								<div class="form-group">
									<input  type="text" name="email" class="input-control" placeholder="Email">
								</div>
								<div class="form-group">
									<input  type="text" name="phone" class="input-control" placeholder="Phone no.">
								</div>
							</div>
							<div class="form-group">
								<textarea    class="input-control" placeholder="Your message" name="message"></textarea>
							</div>
							<div class="form-group">
								<button type="submit" class="btn btn-alter btn-border btn-border-brown" >Submit</button>
							</div>
						</form>
					</div><!-- /.form-contact -->
					<div class="address-wrap clearfix">
						<div class="address-info">
							<ul>
								<li><i class="icon icon-map-white"></i>Delo Road, Dr. Grahams Homes Block B, Kalimpong, WB - 734301</li>
								<li><i class="icon icon-phone"></i>+91 76020 29626</li>
								<li><i class="icon icon-mail"></i>contact@aromahousekpg.com</li>
							</ul>
						</div>
						<div class="address-map">
							<div>
								<i class="icon icon-map-brown"></i>
								<a class="popup-gmaps" href="https://www.google.com/maps/place/Kalimpong+Broadband/@27.0880823,88.520325,17z/data=!3m1!4b1!4m5!3m4!1s0x39e41fbc62137019:0x31d0560042ad0f95!8m2!3d27.0880775!4d88.5225137">See Map</a>
							</div> 
						</div>
					</div><!-- /.address-wrap -->
				</div>
			</div>
		</section><!-- /.md-wrapper  -->
	{{ end }}

{{ end }}

{{ define "js" }}
<script type="text/javascript" src="/js/jquery.form.js"></script>
<script type="text/javascript" src="/js/jquery.validate.min.js"></script>
<script type="text/javascript" src="/js/jquery.magnific-popup.min.js"></script>
<script type="text/javascript" src="/js/jquery.magnific-popup.map.js"></script>
<script type="text/javascript">
	$(function() {
		"use strict";
		$("#arrival-date, #departure-date").datepicker({
			// showOn: "button",
			// buttonImage: "images/calendar.gif",
			// buttonImageOnly: true
		});

		// Script for Popup Map Address
 	// 	$('.popup-gmaps').magnificPopup({
		// 	// disableOn: 700,
		// 	type: 'iframe',
		// 	mainClass: 'md-map',
		// 	removalDelay: 160,
		// 	preloader: false,
		// 	fixedContentPos: false
		// });

	$('.popup-gmaps').magnificPopupMap();

		if($("#contact-form").length > 0){
	      // Validate the contact form
	      $('#contact-form').validate({
	        // Add requirements to each of the fields
	        rules: {
	          name: {
	            required: true,
	            minlength: 2
	          },
	          email: {
	            required: true,
	            email: true
	          },
	          message: {
	            required: true,
	            minlength: 10
	          }
	        },

	        // Specify what error messages to display
	        // when the user does something horrid
	        messages: {
	          name: {
	            required: "Please enter your name.",
	            minlength: $.format("At least {0} characters required.")
	          },
	          email: {
	            required: "Please enter your email.",
	            email: "Please enter a valid email."
	          },
	          message: {
	            required: "Please enter a message.",
	            minlength: $.format("At least {0} characters required.")
	          }
	        },

	        // Use Ajax to send everything to processForm.php
	        submitHandler: function(form) {
	          $("#submit-contact").html("Sending...");
	          $(form).ajaxSubmit({
	            success: function(responseText, statusText, xhr, $form) {
	              $("#contact-content").slideUp(600, function() {
					$("#contact-content").html(responseText).slideDown(600);
				  });
				  $("#submit-contact").html("Submit");
	            }
	          });
	          return false;
	        }
	      });
	    }
	});

 	</script>

{{ end }}