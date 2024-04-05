<script>
	import '../app.css';
</script>

<div class="flex flex-row">
	<div class="">
		<nav>
			<a href="/">Home</a>
			<a href="/fav">Favorites</a>
		</nav>
	</div>
	<div class="">
		<slot />
	</div>
</div>
