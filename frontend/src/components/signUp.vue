<template>
<div class="view">
    <h1>註 冊</h1>
    <div class="box">
        <div class="account">
            <h4>帳號</h4>
            <el-input v-model="account" maxlength = "14" size="large" placeholder="請輸入您要創建的帳號" />
        </div>
        <div class="password">
            <h4>密碼</h4>
            <el-input v-model="password" maxlength = "14" type="password" size="large" placeholder="請輸入您要創建的密碼" />
        </div>
        <div class="name">
            <h4>姓名</h4>
            <el-input v-model="name" maxlength = "30" size="large" placeholder="請輸入您的姓名" />
        </div>
        <div class="phone">
            <h4>手機號碼</h4>
            <el-input v-model="phone" maxlength = "12" size="large" placeholder="請輸入您的手機號碼" />
        </div>
        <el-button @click="signUp" class="signUpBtn" type="warning">註冊</el-button>
    </div>

</div>
</template>
    <script setup>
    import { useRouter} from "vue-router";
    import { onMounted, ref, reactive } from 'vue';
    import axios from  'axios';
    const router = useRouter();
    const account = ref('');
    const password = ref('');
    const name = ref('');
    const phone = ref('');
    // //-------------------------------
    // const Home = { template: '<div>Home</div>' }
    // const routes = [
    //     { path: '/product.vue'},

    // ]
    // const router = VueRouter.createRouter({
    // // 4. Provide the history implementation to use. We are using the hash history for simplicity here.
    //     history: VueRouter.createWebHashHistory(),
    //     routes, // short for `routes: routes`
    // })
    // //------------------------------
    function signUp(){
         axios
         .post('https://kinny-shop-backend-ezdtsdzl7q-uc.a.run.app/signUp', {
            account: account.value,
            password: password.value,
            phone: phone.value
        })
         .then(function (response){
            if(response.data.stas == "succeeded"){
                alert("註冊成功")
                router.push({ name: "Login" });
            } else{
                alert(response.data.res)
            }
         })
         .catch(function (error) {
             alert(error);
        });
    }

</script>
<style>
    body{
    background-color:#ffa;
    }
    .view{
        margin-top: 90px;
        text-align: center;
    }
    .box{
        display: inline-block;
        width: 360px;
        height: 460px;
        border: 2px #ddd solid;
        border-radius: 10px;
        background-color: #eee;
    }
    h1{
        font-family: sans-serif;
        font-weight: bolder;
        margin-bottom: 15px;
        color: rgb(128, 128, 128);
    }
    h4{
        font-family: monospace;
        font-size: 16px;
        color:#aaa;
        text-align: left;
        margin-top: 30px;
        margin-bottom: 5px;
    }
    .account ::-webkit-input-placeholder { 
        font-family: monospace;
        font-size: 12px;
    }
    .name ::-webkit-input-placeholder { 
        font-family: monospace;
        font-size: 12px;
    }
    .password ::-webkit-input-placeholder { 
        font-family: monospace;
        font-size: 12px;
    }
    .phone ::-webkit-input-placeholder { 
        font-family: monospace;
        font-size: 12px;
    }
    .account{
        display: inline-block;
        height: 50px;
        width: 280px;
    }
    .password{
        display: inline-block;
        height: 50px;
        width: 280px;
    }
    .name{
        display: inline-block;
        height: 50px;
        width: 280px;
    }
    .phone{
        display: inline-block;
        height: 50px;
        width: 280px;
    }
    .signUpBtn{
        margin-top: 40px;
        height: 50px;
        width: 280px;
        border-radius: 25px;
        font-family: monospace;
        font-weight: bolder;
    }
</style>