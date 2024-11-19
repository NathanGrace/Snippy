<script lang="ts">
	import '../app.css';
	let { children } = $props();
</script>


<div class="navbar bg-base-100">
	<div class="flex-1">
		<a href="./" class="btn btn-ghost text-xl">Snippy</a>
	</div>
	<div class="flex-none">
		<ul class="menu menu-horizontal px-1">
			<li><a href="/calendar">Calendar</a></li>
			<li>
				<details>
					<summary>Parent</summary>
					<ul class="bg-base-100 rounded-t-none p-2">
						<li><a href="./">Link 1</a></li>
						<li><a href="./">Link 2</a></li>
					</ul>
				</details>
			</li>
		</ul>
	</div>
</div>

{@render children()}
