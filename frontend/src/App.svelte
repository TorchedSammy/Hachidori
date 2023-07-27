<script lang="ts">
  import svelteLogo from './assets/svelte.svg'
  import viteLogo from '/vite.svg'
  import Counter from './lib/Counter.svelte'
  import { Navbar, NavBrand, NavLi, NavUl, NavHamburger, Search } from 'flowbite-svelte'

  let value = '';
</script>

<Navbar let:hidden let:toggle>
	<NavBrand href="/">
	  <img
		src="{svelteLogo}"
		class="mr-3 h-6 sm:h-9"
		alt="Logo"
	  />
	  <span class="self-center whitespace-nowrap text-xl font-semibold dark:text-white">
		Hachidori
	  </span>
	</NavBrand>
	<NavHamburger on:click={toggle} />
	<NavUl {hidden}>
	  <NavLi href="/" active={true}>Home</NavLi>
	  <NavLi href="/about">About</NavLi>
	  <NavLi href="/services">Services</NavLi>
	  <NavLi href="/pricing">Pricing</NavLi>
	  <NavLi href="/contact">Contact</NavLi>
	</NavUl>
</Navbar>
<main>

<Search bind:value />
<h1>blud is searching {value}</h1>

  <h1>Vite + Svelte</h1>

  <div class="card">
	<Counter />
  </div>

  <p>
	Check out <a href="https://github.com/sveltejs/kit#readme" target="_blank" rel="noreferrer">SvelteKit</a>, the official Svelte app framework powered by Vite!
  </p>

  <p class="read-the-docs">
	Click on the Vite and Svelte logos to learn more
  </p>
</main>

<style>
  .read-the-docs {
	color: #888;
  }
</style>
