<script lang="ts">
	let { data } =$props();
</script>

<h1>Login Page</h1>

<form action="?/login" method="post">
	<input type="text" name="username" placeholder="Username" />
	<input type="password" name="password" placeholder="Password" />
	<button type="submit">Login</button>
</form>

{#if data.user}
	<p>Hello {data.user.username}</p>
{/if}