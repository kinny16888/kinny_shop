<template>
<div class="view">
    <h1>登 入</h1>
    <div class="box">
        <div class="account">
            <h4>帳號</h4>
            <el-input v-model="account" maxlength = "14" size="large" placeholder="請輸入您的帳號" />
        </div>
        <div class="password">
            <h4>密碼</h4>
            <el-input v-model="password" maxlength = "14" type="password" size="large" placeholder="請輸入您的密碼" />
        </div>
        <button class="astext" @click="SignUp">還沒有帳號嗎?點擊前往註冊</button>
        <el-button @click="login" class="signBtn" type="warning">登入</el-button>
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
    function login(){
         axios
         .post('https://kinny-shop-backend-ezdtsdzl7q-uc.a.run.app/login', {
            account: account.value,
            password: password.value
        })
         .then(function (response){
            if(response.data.stas == "succeeded"){
                localStorage.account = account.value
                localStorage.jwtToken = response.data.res
                alert("登入成功")
                if (response.data.permission == 1){
                    localStorage.permission = response.data.permission
                    router.push({ name: "Manager" });
                }else{
                    router.push({ name: "Product" });
                }
            } else{
                alert(response.data.res)
            }
         })
         .catch(function (error) {
             alert(error);
        });
    }
    function SignUp(){
         router.push({ name: "SignUp" });
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
        height: 420px;
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
    .account{
        display: inline-block;
        height: 50px;
        width: 280px;
    }
    .account ::-webkit-input-placeholder { 
        font-family: monospace;
        font-size: 12px;
    }
    .password ::-webkit-input-placeholder { 
        font-family: monospace;
        font-size: 12px;
    }
    .password{
        display: inline-block;
        height: 50px;
        width: 280px;
    }
    .signBtn{
        height: 50px;
        width: 280px;
        border-radius: 25px;
        font-family: monospace;
        font-weight: bolder;
        margin-top: 10px;
    }
    .astext {
        color: blue;
        font-weight: bolder;
        font-size: 14px;
        background: none;
        border: none;
        margin: 20px 80px 0 0;
        cursor: pointer;
    }
</style>