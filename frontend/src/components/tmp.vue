<template>
</template>
    <script setup>
    import { useRouter} from "vue-router";
    import { onMounted, ref, reactive } from 'vue';
    import axios from  'axios';
    const router = useRouter();
    onMounted(() => {
        if(localStorage.permission == "1"){
            router.push({ name: "Manager" });
        }else{
            router.push({ name: "ShoppingCart" });
        }
        
    })
</script>
<style>

</style>