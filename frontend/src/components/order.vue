<template>
<div class="container">
    <div class="title">
        <el-button @click ="go()"><el-icon :size="22"><Back /></el-icon>Back</el-button>
        <h3>訂單詳情</h3>
    </div>
        <div class="item_header">
            <div class="uid">會員帳號</div>
            <div class="name">會員名稱</div>
            <div class="adderss">地址</div> 
            <div class="date">日期</div>
            <div class="sendMark">訂單狀態</div>
            <div class="check_order">查看訂單</div>
        </div>
        <div class="item_container" v-for="(item, index) in itemList" :key="item.id" >
            <div class="item_header item_body">
                <div class="uid">
                    {{item.uid}}
                </div>
                <div class="name">
                    {{item.name}}
                </div>
                <div class="address">       
                    {{ item.address }}
                </div> 
                <div class="date">
                    {{item.date}}
                </div>
                <div class="sendMark">
                    <el-button v-if="item.sendMark == true" type="danger" >已出貨</el-button>
                    <el-button v-else type="success" >尚未出貨</el-button>
                </div>
                <div class="check_order">
                        <el-button type="warning" @click="dialogFormVisible = true; detailedOrder(item.pkno);">訂單詳情</el-button>
                </div>
            </div>
        </div>
        <div class="moreForm">
            <el-dialog v-model="dialogFormVisible" title="訂單詳情">
                <el-form class="tidyForm" :model="form" v-for="(item, index) in orderList" :key="item.index">
                    <el-form-item class="formItem" label="產品名稱：" :label-width="formLabelWidth">
                        <span>{{ item.pname }}</span>
                    </el-form-item>
                    <!-- <el-form-item label="售價：" :label-width="formLabelWidth">
                        <span>{{ detailedItem.price }}</span>
                    </el-form-item> -->
                    <el-form-item class="formItem" label="購買數量：" :label-width="formLabelWidth">
                        <span>{{ item.quantity }}</span>
                    </el-form-item> 
                </el-form>
                <template #footer>
                <span class="dialog-footer">
                    <el-button type="primary" @click="dialogFormVisible = false">關閉</el-button>
                </span>
                </template>
            </el-dialog>
        </div>
    </div>
</template>

<script setup lang="ts">
import { ref, reactive, onMounted } from 'vue'
    import axios from  'axios';
    import { useRouter} from "vue-router";

    const router = useRouter();
    let playList = ref();
    let itemList = ref();
    let orderList = ref();
    let dialogFormVisible = ref(false);
    const userInfo = localStorage.account
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
        if(userInfo == null || jwtToken == null){
            router.push({ name: "Product" });
            return
        }
        serachOrder()
    })
    function go(){
        router.push({ name: "Product" });
    }
    function serachOrder(){
        axios
        .get('https://kinny-shop-backend-ezdtsdzl7q-uc.a.run.app/order', {headers:{"Authorization" : `Bearer ${jwtToken}`}} )
        .then(function (response){
            if(response.data.stas=="failed"){
                alert(response.data.res)
                return
            }
            itemList.value = response.data.res;
        })
        .catch(function (error) {
           alert(error);
        });
    }
    function detailedOrder(pkno){
        axios
        .get('https://kinny-shop-backend-ezdtsdzl7q-uc.a.run.app/order/detailedOrder/'+pkno,{headers:{"Authorization" : `Bearer ${jwtToken}`}} )
        //.get('http://localhost:80/select.php')
        .then(function (response){
            console.log(response.data.res.commodity)
            orderList.value = response.data.res.commodity
            console.log(orderList.value)
        })
        .catch(function (error) {
           alert(error);
        });
    }
</script>

<style>
.item_header{
    display: flex;
    width: 1000px;
    margin: 0 auto;
    height: 30px;
    background-color: rgba(99, 99, 99, 0.9);
    border-radius: 5px;
    padding-left: 10px;
}
.item_header div{
    width: 200px;
    color: white;
    line-height: 30px;
}
.item_body{
    margin-top: 1px;
    height: 100px;
    align-items: left;   
}
.tidyForm{
    display: flex;
}
.formItem{
    width: 250px;
}
.tidyEditProduct{
    display: inline-block;
    width: 540px;
}
.title{
    display: flex;
}
.title h3{
    margin-left: 90px;
    font-family: Microsoft JhengHei;
}
</style>