<script setup>
import { DuplicateIcon, ExternalLinkIcon } from '@heroicons/vue/solid'
import clip from 'clipboard'
import Button from './Button.vue';
    const props = defineProps({
        longUrl: String,
        slug: String,
    })
function copy() {
    clip.copy(`www.delongify.xyz/${props.slug}`)
}
</script>

<template>
    <div class="flex flex-col w-full sm:w-2/3 md:w-3/5 lg:w-1/3 items-center py-4 px-6 bg-indigo-600 rounded-lg space-y-2 shadow">
        <span class="w-full text-ellipsis overflow-hidden text-sm opacity-75">{{longUrl}}</span>
        <span class="px-3 py-1 bg-indigo-700 rounded-lg">www.delongify.xyz/{{slug}}</span>
        <span class="flex space-x-2">
            <Button color="indigo">
                <span class="flex items-center space-x-1">
                    <ExternalLinkIcon class="w-5 h-5"></ExternalLinkIcon>
                    <span>Visit</span>
                </span>
            </Button>
            <Button color="fuchsia" v-on:click="copy()">
                <span class="flex items-center space-x-1">
                    <DuplicateIcon class="w-5 h-5"></DuplicateIcon>
                    <span>Copy</span>
                </span>
            </Button>
        </span>
        <span class="text-xs opacity-50">This link will expire in 24 hours.</span>
    </div>
</template>

<style scoped>
    
</style>