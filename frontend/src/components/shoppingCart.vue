<template>

<div class="container">
    <el-button @click ="go()"><el-icon :size="22"><Back /></el-icon>Back</el-button>
    <div class="item_header">
        <div class="item_detail">商品</div>
        <div class="price">單價</div>
        <div class="allcount">數量</div> 
        <div class="amount">小計</div>
        <div class="operate">操作</div>
    </div>
    <div class="item_container" v-for="(item, index) in itemList" :key="item.id">
        <div class="item_header item_body">
            <div class="item_detail">
                <!-- <img v-bind:src="item.imgUrl" alt=""> -->
                <div class="name">{{item.pname}}</div>
            </div>
            <div class="price"><span>$</span>{{item.price}}</div>
            <div class="buyCount">
                <el-button size="default" :disabled="btnLock" type="info" plain @click="quantityControl(index,'-')">-</el-button>
                {{item.quantity}}
                <el-button size="default" :disabled="btnLock" type="info" plain @click="quantityControl(index,'+')">+</el-button>
            </div> 
            <div class="amount">{{item.price * item.quantity}}</div>
            <div class="operate">
                <el-button type="danger" @click="handledelete(item.pid)">刪除</el-button>
            </div>
        </div>
    </div>
    <el-divider></el-divider>
    <div class="totalAmount"> 
        <p>總計:{{totalAmount()}}</p>
        <el-button type="primary" :disabled="submitLock" @click="totalAmountDialogFormVisible = true">送出訂單</el-button>
    </div>  
    <div class="totalAmountForm">
            <el-dialog v-model="totalAmountDialogFormVisible" title="訂單內容">
                <el-form :model="form">
                    <el-form-item label="" :label-width="formLabelWidth"  v-for="(item, index) in itemList" :key="item.id">
                        <span>產品名稱：{{ item.pname }}　　</span>
                        <span>購買數量：{{ item.quantity }}　　</span>
                    </el-form-item>
                    <el-form-item label="" :label-width="formLabelWidth">
                        <span>總計：{{totalAmount()}}</span>
                    </el-form-item>
                    <el-form-item label="地址：" class="inputAddress" :label-width="formLabelWidth">
                        <el-input v-model="adderss" maxlength = "100" size="large" placeholder="請輸入您的地址" />
                    </el-form-item>
                </el-form>
                <template #footer>
                <span class="dialog-footer">
                    <el-button @click="totalAmountDialogFormVisible = false">取消</el-button>
                    <el-button type="primary" :disabled="beLock" @click="submitOrder">提交</el-button>
                    <el-divider></el-divider>
                </span>
                </template>
            </el-dialog>
        </div>
</div>
</template>
<script setup>
    import axios from  'axios';
    import { ref, reactive, onMounted } from 'vue'
    const userInfo = localStorage.account
    const jwtToken = localStorage.jwtToken
    let itemList = ref('')
    const buyCount = ref();
    const adderss = ref();
    let totalAmountDialogFormVisible = ref(false);
    buyCount.value = 0;
    onMounted(() => {
        if(localStorage.permission == 1){
            router.push({ name: "Manager" });
            return
        }
        if(userInfo == null || jwtToken == null){
            router.push({ name: "Product" });
            return
        }
        showShoppingCart();
    })
    const beLock =  ref(false);
    const btnLock = ref(false);
    const submitLock = ref(false);
    function submitOrder(){
        // console.log(itemList.value)
        beLock.value = true;
        axios
        .post('https://kinny-shop-backend-ezdtsdzl7q-uc.a.run.app/order', 
            {
                amount: totalAmount(),
                address : adderss.value,
                commodity:itemList.value
            }, 
            {
                headers:{"Authorization":`Bearer ${jwtToken}`}
            }
        )
        .then(function (response){
            beLock.value = false;
            totalAmountDialogFormVisible.value = false;
            var show = "剩餘數量: \n"
            if(response.data.stas == "-1"){
                alert("商品庫存不夠")
                for(var i = 0;i<response.data.res.length;i++){
                    show += response.data.res[i].Name
                    show += " : "
                    show += response.data.res[i].Stock
                    show += "\n"
                }
                alert(show)
                return
            }else{
                alert(response.data.res)
            }
            router.push({ name: "Tmp" });
        })
        .catch(function (error) {
            alert(error);
        });
    }
    function totalAmount(){
        let total = 0;
        for(let i = 0;i<itemList.value.length;i++){
            total += parseInt(itemList.value[i].price) * parseInt(itemList.value[i].quantity);
        }
        return total
    }
    function showShoppingCart() {
        //alert(account)
        axios
        .get('https://kinny-shop-backend-ezdtsdzl7q-uc.a.run.app/shoppingCart', {headers:{"Authorization" : `Bearer ${jwtToken}`}} )
        .then(function (response){
            if(response.data.res == null){
                submitLock.value = true
                alert("購物車是空的")
                return
            }
            itemList.value = response.data.res;
            //const files = require.context('../assets/1', false, /.jpg$/).keys();
            console.log(itemList.value[0])
            // buyCount.value = 
        })
        .catch(function (error) {
            alert(error);
        });
    }
    
    function quantityControl(index,sign){
        btnLock.value=true
        if(sign == '+'){
            if(itemList.value[index].quantity<itemList.value[index].stock){
                axios
                .put('https://kinny-shop-backend-ezdtsdzl7q-uc.a.run.app/shoppingCart/'+sign, 
                    {
                        pid: itemList.value[index].pid
                    }, 
                    {
                        headers:{"Authorization":`Bearer ${jwtToken}`}
                    }
                )
                .then(function (response){
                    showShoppingCart()
                    router.push({ name: "Tmp" });
                })
                .catch(function (error) {
                    alert(error);
                });
            }else{
                return
            }
        }else{
            if(itemList.value[index].quantity>1){
                axios
                .put('https://kinny-shop-backend-ezdtsdzl7q-uc.a.run.app/shoppingCart/'+sign, 
                    {
                        pid: itemList.value[index].pid
                    }, 
                    {
                        headers:{"Authorization":`Bearer ${jwtToken}`}
                    }
                )
                .then(function (response){
                    showShoppingCart()
                    router.push({ name: "Tmp" });
                })
                .catch(function (error) {
                    alert(error);
                });
            }else{
                return
            }
        }
    }
    function go(){
        router.push({ name: "Product" });
    }
    function handledelete(pkno) {
        //this.itemList.splice(index,1);
        var check = confirm("你確定要刪除嗎?");
        var tmp = String(pkno)
        if(check){
            axios
            .delete('https://kinny-shop-backend-ezdtsdzl7q-uc.a.run.app/shoppingCart/'+tmp,{headers:{"Authorization" : `Bearer ${jwtToken}`}})
            .then(function (response){
                alert(response.data.res)
                showShoppingCart();
                router.push({ name: "Tmp" });
            })
            .catch(function (error) {
                alert(error);
            });
        }else{
            return
        }
        
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
    font-size: 16px;
}
.item_header div{
    width: 200px;
    color: white;
    line-height: 30px;
}
.item_header .item_detail{
    width: 50%;
}
.item_body{
    margin-top: 1px;
    height: 100px;
    align-items: center;  
    font-size: 16px;  
}
.item_detail img{
    width: 80px;
    height: 80px;
    border-radius: 3px;
    /* margin-top: 10px; */
    float: left;
}
.item_detail .name{
    margin-left: 100px;
    margin-top: 20px;
} 
.price{
    text-align: center;
}
.amount{
    text-align: center;
}
.allcount{
    text-align: center;
}
.operate{
    text-align: center;
}
.totalAmount{
    width: 1150px;
    display: flex;
    font-family: Microsoft JhengHei;
    font-weight: bold;
    align-items: center;
    justify-content:flex-end;
}
.totalAmount p{
    margin: 0 30px;
}
.buyCount {
    display: flex;
    align-items: center;
    justify-content: center;
    font-size: 16px;
}
.buyCount button{
    margin: 0 8px;
}
.inputAddress{
    display: flex;
    height: 50px;
    width: 280px;
}
.inputAddress ::-webkit-input-placeholder { 
    font-family: monospace;
    font-size: 12px;
}
</style>