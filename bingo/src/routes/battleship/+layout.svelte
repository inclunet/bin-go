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
	header h1 { width:clamp(30rem, 65vw, 60rem); font-size:6.2rem; margin:2rem auto 0; background:var(--secondary-color, rgba(43,127,244,.67)); color:var(--quinary-color,#000); text-shadow:-.1rem -.1rem .3rem #fff, -.1rem .1rem .3rem #fff, .1rem -.1rem 0 #fff, .1rem 0 0 #fff; border-radius:1rem; text-align:center; }
	.top-header_container { display:flex; margin:2.4rem 0 3rem 0; }
	.content { flex:1 0 auto; width:100%; max-width:none; margin:0; padding:0 1rem 2rem; box-sizing:border-box; }
	a { color:#fff; }
	@media (max-width:794px){ header h1 { font-size:5.2rem; } header img { width:100%; } }
	@media (max-width:768px){
		.content { padding-left:0.5rem; padding-right:0.5rem; }
		header h1 { max-width:100%; box-sizing:border-box; }
	}
	@media (max-width:685px){ header { padding:2.5rem 2.2rem 1rem; } header h1 { font-size:4.6rem; width:100%; } }
	@media (max-width:640px){ header { padding:2.3rem 1.4rem 1rem; } header h1 { font-size:4.2rem; } }
	@media (max-width:571px){ header { justify-content:space-between; } header h1 { width:100%; margin:0 auto; padding:.4rem .6rem; font-size:3.8rem; } header img { width:78%; } }
	@media (max-width:450px){ header { flex-direction:row; padding:.3rem 1rem; align-items:center; background-image:none; } .top-header_container { margin:0; } header a { display:none; } header h1 { width:100%; font-size:3.4rem; margin:0; } }
	@media (max-width:480px){
		.content { padding-left:0.35rem; padding-right:0.35rem; }
	}
</style>
