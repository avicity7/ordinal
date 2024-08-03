<script>
	import Button from "$lib/components/ui/button/button.svelte";

</script>
<div class="p-24">
  <h1 class="text-6xl font-medium">Yet another LMS site.</h1>
  <h2 class="text-3xl mt-8">Go use Google Classroom or something.</h2>

  <div class="flex flex-row w-full justify-center mt-40">
    <div class="flex flex-col items-center">
      <h4 class="text-3xl">Or, you can stay here;</h4>
      <h3 class="text-4xl font-medium mt-4">we've got the best coding tutorials.</h3>
    </div>
  </div>

  <div class="flex flex-row w-full justify-center mt-40">
    <div class="flex flex-col items-center">
      <h4 class="text-3xl mb-2"><span class="font-medium">3</span> coding languages</h4>
      <h4 class="text-3xl mb-2"><span class="font-medium">100+</span> videos</h4>
      <h4 class="text-3xl mb-2"><span class="font-medium">300+</span> questions</h4>
    </div>
  </div>

  <div class="flex flex-row w-full justify-center mt-40">
    <div class="flex flex-col items-center">
      <h3 class="text-4xl font-medium my-8">Get started now.</h3>
      <Button href="/profile/signup">Sign up</Button>
    </div>
  </div>
</div>
