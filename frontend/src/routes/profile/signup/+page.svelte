<script lang="ts">
	import { onMount } from 'svelte';

  export let data
  
  onMount(() => {
    if (data.session) {
      window.location.replace('/profile')
    }
  })
</script>

<svelte:head>
  <script src="https://www.google.com/recaptcha/api.js?render=explicit" async defer></script>
	<title>Sign up | SPOH Bus Tracker</title>
</svelte:head>

<div class="w-screen flex flex-col items-center text-zinc-800">
  <div class="max-w-xl mt-12">
    <h1 class="font-bold text-3xl w-full mb-6">Sign Up</h1>
    <form 
      class="flex flex-col"
      method='POST'
      action="?/signUp"
    >
      <div class="font-medium mb-4 mt-2">Name</div>
      <input 
        type="text"
        autocomplete="off"
        name="name"
        class="px-3 border-2 border-zinc-200 rounded-md focus:border-yellow-500 outline-none select-none h-[2rem] mr-2 dark:text-zinc-800"
      />
      <div class="font-medium mb-4 mt-8">Email</div>
      <input 
        type="text"
        autocomplete="off"
        name="email"
        class="px-3 border-2 border-zinc-200 rounded-md focus:border-yellow-500 outline-none select-none h-[2rem] mr-2 dark:text-zinc-800"
      />
      <div class="font-medium mb-4 mt-8">Password</div>
      <input 
        type="password"
        autocomplete="off"
        name="password"
        class="px-3 border-2 border-zinc-200 rounded-md focus:border-yellow-500 outline-none select-none h-[2rem] mr-2 dark:text-zinc-800 mb-8"
      />
      <button class="py-1 bg-primary hover:bg-yellow-700 font-medium text-white rounded-md mt-6 transition-colors" type="submit">Sign up</button>
      <div class="text-zinc-400 font-medium text-sm mx-6 md:mx-24 mt-8 text-center">
        Already have an account?
        <a href="/profile/signin" class="text-blue-600 hover:text-blue-400">
          Sign in
        </a>
      </div>
    </form>
  </div>
</div>
