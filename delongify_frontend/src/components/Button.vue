<script setup>
    import { computed } from 'vue'
    const props = defineProps({
        method: Function,
        color: String
    })
    const classObject = computed(() => ({
        'bg-fuchsia-600 hover:bg-fuchsia-500 active:bg-fuchsia-600': props.color == 'fuchsia',
        'bg-indigo-500 hover:bg-indigo-400 active:bg-indigo-500': props.color == 'indigo',
    }))
</script>
<template>
    <button class="rounded-lg px-3 py-1.5 transition duration-50 shadow" :class="classObject">
        <slot></slot>
    </button>
</template>
<style scoped>

</style>