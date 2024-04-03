<script>
	import '../app.css';
</script>

<div class="flex">
	<div class="w-1/12">
		<nav>
			<a href="/">Home</a>
			<a href="/fav">Favorites</a>
		</nav>
	</div>
	<div class="flex-grow">
		<slot />
	</div>
</div>
