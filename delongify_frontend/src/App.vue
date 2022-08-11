<script setup>
  import { CakeIcon, DotsHorizontalIcon } from '@heroicons/vue/solid'
  import Button from './components/Button.vue'
  import { ref } from 'vue'
  import axios from 'axios'
import DelongifiedLinkCard from './components/DelongifiedLinkCard.vue'
  document.title = "Delongify | Home"

  let slugs = ref([])
  // data
  let isGenerating = ref(false)
  let longUrl = ref("")

  function delongifyUrl(url) {
    isGenerating.value = true
    axios.post("http://delongify.xyz/api/shorten", {
      Url: url,
    })
    .then(function (res) {
      slugs.value.unshift(
        {
          slug: res.data.Slug,
          url: res.data.Url,
        }
      )
      console.log(res)
    })
    .catch(function (err) {
      console.log(err)
    })
    .finally(() => {
      isGenerating.value = false
    })
  }

</script>

<template>
  <div class="flex flex-col w-full min-h-screen items-center p-4 space-y-4 bg-indigo-800 text-white">
    <div class="flex flex-col w-full sm:w-2/3 md:w-3/5 lg:w-1/3 items-center py-8 px-6 bg-indigo-600 rounded-lg space-y-2 shadow">
      <span class="pacifico text-5xl pt-2 pb-8">Delongify</span>
      <span>Enter a url to turn into a <span class="pacifico">delongified</span> url:</span>
      <input type="url" class="w-full px-3 py-1 rounded-lg text-black" v-model="longUrl">
      <span class="flex pt-1">
        <Button v-on:click="delongifyUrl(longUrl)" color="fuchsia">
          <span class="flex space-x-1" v-if="!isGenerating">
            <CakeIcon class="w-5 h-5"></CakeIcon>
            <span>Shorten URL</span>
          </span>
          <span class="flex space-x-1" v-if="isGenerating">
            <DotsHorizontalIcon class="w-5 h-5 animate-pulse"></DotsHorizontalIcon>
          </span>
        </Button>
      </span>
      <span class="text-sm text-center opacity-75 pt-5">By clicking Shorten URL I agree to the <span class="underline">Terms of Service</span>, <span class="underline">Privacy Policy</span>, and Use of Cookies.</span>
    </div>

    <DelongifiedLinkCard v-for="slug in slugs" :key="slug" :long-url="slug.url" :slug="slug.slug"></DelongifiedLinkCard>
  </div>
</template>

<style scoped>
  @import url('https://fonts.googleapis.com/css2?family=Pacifico&display=swap');
  .pacifico {
    font-family: 'Pacifico', cursive;
  }
</style>
