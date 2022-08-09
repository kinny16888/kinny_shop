<template>
    <div class="userButton">
       <el-button type="primary" @click ="goOrder()"><el-icon :size="22"><Avatar /></el-icon></el-button>
    </div>
    <div class="signOut">
        <el-button v-if="account!=null || jwtToken!=null"  type="danger" @click ="signOut()">登出</el-button>
    </div>
    <div class="signButton">
        <label>您好! {{account}} </label>
        <el-button v-if="account==null && jwtToken==null" type="success" @click ="goLogin()">登入</el-button>
        <el-button type="success" v-else @click ="goShoppingCart()"><el-icon :size="22"><ShoppingCart /></el-icon></el-button>
        <!-- <el-button
            v-for="Button in buttons"
            :key="Button.text"
            :type="Button.type"
            link
            >{{ Button.text }}</el-button> -->
    </div>

    <div class="search">
        <div class="searchInput">
            <el-input
                v-model="serachText"
                maxlength = "30"
                class="w-50 m-2"
                size="large"
                placeholder="搜尋"
            />
        </div>
        <el-button @click ="searchProduct()" type="primary" :icon="Search" circle />
    </div>

<div class="tidyCard">
    <div class="Listcard" >
        <el-card class="box-card" v-for="(itemm ,index) in itemList" >
            <template #header>
                <div class="card-header">
                    <!-- <span>{{ itemList[i-1+index].name }}</span> -->
                    <span>{{ itemm.name }}</span>
                </div>
            </template>
            <img
                :src="'https://kinny-shop-backend-ezdtsdzl7q-uc.a.run.app/file?folder='+itemm.pkno+'&fileName='+itemm.filePath[0]"
                :alt="itemm.name"
                class="image"
            />
            <div class="text-item">
                售價：{{ itemm.price }} <br>
                庫存：{{ itemm.stock }} <br>
            </div>
            <el-button type="warning" class="moreButton" @click="dialogFormVisible = true; productDetailed(itemm.pkno);">點我看更多...</el-button>
        </el-card>
        <div class="moreForm">
            <el-dialog v-model="dialogFormVisible" title="商品">
                <el-form :model="form">
                    <el-carousel indicator-position="outside">
                        <el-carousel-item v-for="(item ,index) in detailedItem.filePath" :key="item" class="carousel_img">
                            <img
                                :src="'https://kinny-shop-backend-ezdtsdzl7q-uc.a.run.app/file?folder='+detailedItem.pkno+'&fileName='+item"
                                class="image"
                            />
                        </el-carousel-item>
                    </el-carousel>
                    <el-form-item label="產品名稱：" :label-width="formLabelWidth">
                        <span>{{ detailedItem.name }}</span>
                    </el-form-item>
                    <el-form-item label="售價：" :label-width="formLabelWidth">
                        <span>{{ detailedItem.price }}</span>
                    </el-form-item>
                    <el-form-item label="產品描述：" :label-width="formLabelWidth">
                        <span>{{ detailedItem.introduce }}</span>
                    </el-form-item>
                    <el-form-item label="購買數量：" :label-width="formLabelWidth">
                        <div class="count">
                            <el-button size="default" type="info" plain @click="handleSub(detailedItem.stock)">-</el-button>
                            <el-input readonly = "true" v-model="inputCount"/>
                            <el-button size="default" type="info" plain @click="handlePlus(detailedItem.stock)">+</el-button>
                        </div> 
                    </el-form-item>
                </el-form>
                <template #footer>
                <span class="dialog-footer">
                    <el-button @click="dialogFormVisible = false">取消</el-button>
                    <el-button type="primary" @click="addShoppingCart(detailedItem.pkno)">加入購物車</el-button>
                    <el-divider></el-divider>
                </span>
                </template>
            </el-dialog>
        </div>
    </div>
</div>
</template>

<script setup lang="ts">
    import { ref, reactive, onMounted } from 'vue'
    import { Search } from '@element-plus/icons-vue'
    import axios from  'axios';
    import { useRouter} from "vue-router";
    import { ROOT_PICKER_INJECTION_KEY } from 'element-plus';
    // import { createRequire } from 'module';
    // const require = createRequire(import.meta.url);
    const router = useRouter();
    const serachText = ref('');
    const num1 = ref(1);
    // const buttons = [
    //     { type: '', text: '登入' },
    //     { type: '', text: '註冊' }
    // ] as const
    // let oneCardName = ref('');
    // let listItem = ref('');
    // const itemPrice = ref('');
    // const itemIntroduce = ref('');
    // const itemStock = ref('');
    interface Product{
        pkno:number,
        name:string,
        price:number,
        introduce:string,
        stock:number,
        filePath?:string[]
    }
    let tmp : Product={
        pkno:0,
        name:'',
        price:0,
        introduce:'',
        stock:0,
        // filePath:null
    };
    // let itemList = ref([tmp]);
    let itemList = ref();
    let detailedItem = ref(tmp);
    let dialogFormVisible = ref(false);
    let account = localStorage.account;
    let jwtToken = localStorage.jwtToken;
    const formLabelWidth = '140px';
    const form = reactive({
        name: '',
        region: '',
        date1: '',
        date2: '',
        delivery: false,
        type: [],
        resource: '',
        desc: '',
    })
    onMounted(() => {
        if(localStorage.permission == 1){
            router.push({ name: "Manager" });
            return
        }
        axios
        .get('https://kinny-shop-backend-ezdtsdzl7q-uc.a.run.app/product/recommend')
        //.get('http://localhost:80/select.php')
        .then(function (response){
            //console.log(JSON.stringify(response.data.res))
            if(response.data.stas=="failed"){
                alert(response.data.res)
                return
            }
            itemList.value = response.data.res;
            console.log(response.data.res)
        })
        .catch(function (error) {
           alert(error);
        });
    })
    const inputCount = ref();
    inputCount.value = 0;
    function signOut(){
        localStorage.clear()
        location.reload();
    }
    function handlePlus(stock) {
        if(inputCount.value < stock){
            inputCount.value++;
        }
    }
    
    function handleSub(item) {
        if(inputCount.value > 0){
            inputCount.value--; 
        }
    }
    function searchProduct(){
        var url = ""
        if(serachText.value == ""){
            url = "https://kinny-shop-backend-ezdtsdzl7q-uc.a.run.app/product/recommend"
        }else{
            url = "https://kinny-shop-backend-ezdtsdzl7q-uc.a.run.app/product/"+serachText.value
        }
        axios
        .get(url)
        //.get('http://localhost:80/select.php')
        .then(function (response){
            if(response.data.stas=="failed"){
                alert(response.data.res)
                return
            }
            itemList.value = response.data.res;
            // console.log(itemList.value); 
        })
        .catch(function (error) {
           alert(error);
        });
    }
    let num2 = '';
    function productDetailed(pkno) {
        axios
        .get('https://kinny-shop-backend-ezdtsdzl7q-uc.a.run.app/product/detailed/'+pkno)
        //.get('http://localhost:80/detailed.php?pkno='+pkno)
        .then(function (response){
            inputCount.value = 0
            detailedItem.value = response.data.res;
            //const files = require.context('../assets/1', false, /.jpg$/).keys();
            console.log(detailedItem.value.stock);
            console.log(isNaN(detailedItem.value.stock));
            num2 = String(detailedItem.value.stock);
        })
        .catch(function (error) {
           alert(error);
        });
    }
    //const num = Number(detailedItem.value.stock);
    //console.log('num', num);
    function goLogin(){
        router.push({ name: "Login" });
    }
    function goOrder(){
        router.push({ name: "Order" });
    }
    function goShoppingCart(){
        router.push({ name: "ShoppingCart" });
    }
    function addShoppingCart(pkno){
        axios
        .post('https://kinny-shop-backend-ezdtsdzl7q-uc.a.run.app/shoppingCart', 
            {
                pid: pkno,
                quantity : inputCount.value
            }, 
            {
                headers:{"Authorization":`Bearer ${jwtToken}`}
            }
        )
        .then(function (response){
            alert(response.data.res)
        })
        .catch(function (error) {
            if (error.response.data.stas == "failed"){
                alert("請先登入!!")
                router.push({ name: "Login" });
                return
            }
            alert(error);
        });
    }

</script>

<style>
.signButton{
    display: inline-block;
    position: absolute;
    top: 30px;
    right: 100px;
}
.signOut{
    display: inline-block;
    position: absolute;
    top: 30px;
    right: 30px;
}
.signButton label{
    margin-right: 10px;
}
.userButton{
    display: inline-block;
    position: absolute;
    top: 30px;
    left: 40px;
}
.search{
    display: flex;
    justify-content: center; 
    align-items: center; 
}
.searchInput{
    display: inline-block;
    width: 580px;
    margin: 10px 10px;
}

.card-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.text {
  font-size: 14px;
}

.item {
  margin-bottom: 18px;
}

.box-card {
  display: inline-block;
  width: 300px;
  margin-top: 50px;
  margin-left: 25px;
  position: relative;
}
.moreButton{
    position: absolute;
    bottom: 20px;
    right: 20px;
}
.tidyCard{
    display: flex;
    flex-direction: row;
}
.box-card img{
    display:block; 
    margin:auto;
    margin-bottom: 30px;
}
.carousel_img .img{
    display:block; 
    margin:auto;
    margin-bottom: 30px;
}
.count{
    display: flex;
    align-items: center;
}
.count button{
    margin: 0 1px;
}
.count input{
    width: 50px;
}
.image{
    width: 200px;
    height: 200px;
}
</style>