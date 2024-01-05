<script>
	import '../app.css';
	import Chat from '../components/chat.svelte';
	import Footer from '../components/footer.svelte';
	import Header from '../components/header.svelte';
</script>

<Header />

<div class="m-[8vw] xl:m-[4vw]">
	<slot />
</div>

<Footer />
<Chat />
