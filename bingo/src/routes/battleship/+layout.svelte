<script>
	import Footer from "$lib/Footer.svelte";
	import ItemSocial from "$lib/menu/ItemSocial.svelte";
	import MenuHeader from "$lib/menu/MenuHeader.svelte";
	import MenuItem from "$lib/menu/MenuItem.svelte";
	import MenuSocial from "$lib/menu/MenuSocial.svelte";
</script>

<header>
	<div class="top-header_container">
		<a href="/"><img src="/img/logo.png" alt="Inclunet" /></a>
		<h1 class="text-center">Batalha Naval</h1>
	</div>
	<div>
		<MenuHeader title="Batalha Naval">
			<MenuItem href="/">Portal Jogos</MenuItem>
			<MenuItem href="/battleship">Início</MenuItem>
			<MenuItem href="/battleship">Jogar</MenuItem>
			<MenuItem href="/battleship" mobile={true}>Qualquer</MenuItem>
			<MenuSocial>
				<ItemSocial href="https://www.facebook.com/CanalInclunet" src="/img/facebook.png" alt="Facebook" />
				<ItemSocial href="https://www.youtube.com/@Inclunet" src="/img/youtube.png" alt="YouTube" />
				<ItemSocial href="https://www.instagram.com/inclunet/" src="/img/instagram.png" alt="Instagram" />
				<ItemSocial href="https://www.linkedin.com/company/inclunet/" src="/img/linkedin.png" alt="Linkedin" />
				<ItemSocial href="https://www.twitter.com/inclunet" src="/img/twitter.png" alt="Twitter" />
			</MenuSocial>
		</MenuHeader>
	</div>
</header>
<main class="content">
	<slot></slot>
</main>
<Footer />

<style>
	:root { font-size:62.5%; }
	header { display:flex; flex-direction:column; background-color: var(--primary-color, #2b7ef4); background-image:url('/img/fundo.png'); background-repeat:repeat; background-size:50%; }
	header h1 { width:50rem; font-size:6.2rem; margin-left:12%; background:var(--secondary-color, rgba(43,127,244,.67)); color:var(--quinary-color,#000); text-shadow:-.1rem -.1rem .3rem #fff, -.1rem .1rem .3rem #fff, .1rem -.1rem 0 #fff, .1rem 0 0 #fff; border-radius:1rem; }
	.top-header_container { display:flex; margin:2.4rem 0 3rem 0; }
	.content { flex:1 0 auto; }
	a { color:#fff; }
	@media (max-width:794px){ header h1 { font-size:5.4rem; width:42rem; } header img { width:100%; } }
	@media (max-width:685px){ header { padding:2.5rem 2.5rem 1rem; } header h1 { font-size:4.8rem; } }
	@media (max-width:640px){ header { padding:2.3rem 1.5rem 1rem; } }
	@media (max-width:571px){ header { justify-content:space-between; } header h1 { width:32rem; margin:0; padding:0; } header img { width:78%; } }
	@media (max-width:450px){ header { flex-direction:row; padding:.3rem 1rem; align-items:center; background-image:none; } .top-header_container { margin:0; } header a { display:none; } header h1 { width:100%; font-size:3.6rem; } }
</style>
