<template>
<div class="managerSignOut">
    <el-button v-if="userInfo!=null || jwtToken!=null"  type="danger" @click ="managerSignOut()">登出</el-button>
</div>
  <el-tabs v-model="activeName" class="demo-tabs" @tab-click="handleClick">
      <el-tab-pane label="商品管理" name="first">
        <el-button type="primary" @click="dialogInsertProduct = true; clearForm();">新增商品</el-button>
        <div class="container">
            <div class="item_header">
                <div class="productName">商品名稱</div>
                <div class="productPrice">價格</div>
                <div class="productStock">庫存</div>
                <div class="productSendMark">操作</div>
                <div class="productCheck_order">刪除</div>
            </div>
            <div class="item_container" v-for="(item, index) in productList" :key="item.id" >
                <div class="item_header item_body">
                    <div class="productName">
                        {{item.name}}
                    </div>
                    <div class="productPrice">
                        {{item.price}}
                    </div>
                    <div class="productStock">
                        {{item.stock}}
                    </div>
                    <div class="productSendMark">
                         <el-button type="success" @click="dialogEditProduct = true; productDetailed(item.pkno);">編輯</el-button>
                    </div>
                    <div class="productCheck_order">
                         <el-button type="danger" @click="deleteProduct(item.pkno)">刪除</el-button>
                    </div>
                </div>
            </div>
            <div class="insertPoduct"> 
                <el-dialog v-model="dialogInsertProduct" title="新增商品">
                    <el-form class="tidyEditProduct" ref="editNewProduct">
                        <el-form-item class="editProductItem" label="商品名稱" :label-width="formLabelWidth">
                            <el-input v-model="inputInsertName" maxlength="30"></el-input>
                        </el-form-item>
                        <el-form-item class="editProductItem" label="價格" :label-width="formLabelWidth">
                           <el-input v-model="inputInsertPrice" maxlength="8" placeholder="" oninput="value=value.replace(/[^\d]/g,'')"/>
                        </el-form-item> 
                        <el-form-item class="editProductItem" label="庫存" :label-width="formLabelWidth">
                           <el-input v-model="inputInsertStock" maxlength="8" placeholder="" oninput="value=value.replace(/[^\d]/g,'')"/>
                        </el-form-item> 
                        <el-form-item class="editProductItem" label="商品介紹" :label-width="formLabelWidth">
                           <el-input type="textarea" maxlength="100" rows="5" v-model="inputInsertIntroduce" placeholder="" />
                        </el-form-item>
                        <p>第一張圖為預覽圖片!</p>
                        <el-form-item class="editProductItem" label="新增圖片" :label-width="formLabelWidth">
                           <el-upload ref="upload" action="#" list-type="picture-card" :auto-upload="false" :on-change="onfile" accept=".jpg">
                                <el-icon><Plus /></el-icon>
                                <template #file="{ file }">
                                    <div class="flexImg">
                                        <img class="el-upload-list__item-thumbnail" :src="file.url" alt="" />
                                        <span class="el-upload-list__item-actions">
                                            <span
                                                class="el-upload-list__item-preview"
                                                @click="handlePictureCardPreview(file)">
                                                <el-icon><zoom-in /></el-icon>
                                            </span>
                                        </span>
                                    </div>
                                </template>
                            </el-upload>
                            <el-dialog v-model="dialogVisibleImg">
                                <img w-full :src="dialogForImageUrl" alt="Preview Image" />
                            </el-dialog>
                            <!-- <input type="file" @change="onfile" accept="image/png, image/jpeg">
                            <img v-if="testImg" :src="testImg" width="200"/> -->
                        </el-form-item> 
                    </el-form>
                    <template #footer>
                    <span class="dialog-footer">
                        <el-button @click="dialogInsertProduct = false; reload();">取消</el-button>
                        <el-button type="primary" :disabled="isLock" @click="insertProduct()">提交</el-button>
                    </span>
                    </template>
                </el-dialog>
            </div>
             <div class="editProduct">
                <el-dialog v-model="dialogEditProduct" title="編輯商品">
                    <el-form class="tidyEditProduct">
                        <el-form-item class="editProductItem" label="商品編號" :label-width="formLabelWidth">
                            <el-input v-model="inputPkno" readonly = "true" maxlength="30"></el-input>
                        </el-form-item>
                        <el-form-item class="editProductItem" label="商品名稱" :label-width="formLabelWidth">
                            <el-input v-model="inputName" maxlength="8"></el-input>
                        </el-form-item>
                        <el-form-item class="editProductItem" label="價格" :label-width="formLabelWidth">
                           <el-input v-model="inputPrice" maxlength="8" placeholder="" oninput="value=value.replace(/[^\d]/g,'')"/>
                        </el-form-item> 
                        <el-form-item class="editProductItem" label="庫存" :label-width="formLabelWidth">
                           <el-input v-model="inputStock" maxlength="8" placeholder="" oninput="value=value.replace(/[^\d]/g,'')"/>
                        </el-form-item> 
                        <el-form-item class="editProductItem" label="商品介紹" :label-width="formLabelWidth">
                           <el-input type="textarea" maxlength="100" rows="5" v-model="inputIntroduce" placeholder="" />
                        </el-form-item> 
                    </el-form>
                    <template #footer>
                    <span class="dialog-footer">
                        <el-button @click="dialogEditProduct = false">取消</el-button>
                        <el-button type="primary" @click="productSubmit(inputPkno)">提交</el-button>
                    </span>
                    </template>
                </el-dialog>
            </div>
        </div>
    </el-tab-pane>

    <el-tab-pane label="訂單管理" name="second">
        <div class="container">
            <div class="item_header">
                <div class="uid">會員帳號</div>
                <div class="name">會員名稱</div>
                <div class="adderss">地址</div> 
                <div class="date">日期</div>
                <div class="sendMark">操作</div>
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
                        <el-button v-if="item.sendMark == true" type="danger" @click="changeSendMark(item.pkno)">已出貨</el-button>
                        <el-button v-else type="success" @click="changeSendMark(item.pkno)">尚未出貨</el-button>
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
    </el-tab-pane>

    <el-tab-pane label="銷售報表" name="third">
        <div class="flex-item">
            <div class="flex">
                <el-select v-model="selectValue" @change="selectChange" class="m-2" placeholder="請選擇月份" size="large">
                    <el-option
                    v-for="item in selectOptions"
                    :key="item.value"
                    :label="item.label"
                    :value="item.value"
                    />
                </el-select>
                <div class="productDistributed">產品銷售分佈圖</div>
                <div id="chartdiv"></div>
            </div>
            <div class="flex-table">
                <el-table class="table" :data="tableData" stripe style="width: 100%">
                <!-- <el-table-column
                    :prop="index"
                    :label="item"
                    v-for="(item, index) in tableHeader"
                    :key="index"
                    width="180"
                >
                </el-table-column> -->
                <el-table-column prop="key" label="產品" width="180" />
                <el-table-column prop="value" label="銷售金額" width="180" />
                </el-table>
                <div id="barChart"></div>
            </div>
        </div>
    </el-tab-pane>
  </el-tabs>
</template>


<script setup lang="ts">
    import { ref, reactive, onMounted, inject } from 'vue'
    import axios from  'axios';
    import { useRouter} from "vue-router";
    import type { TabsPaneContext , UploadFile, UploadProps } from 'element-plus'
    import { ElMessage } from 'element-plus'
    import { Delete, Download, Plus, ZoomIn } from '@element-plus/icons-vue'
    import * as am4core from "@amcharts/amcharts4/core";
    import * as am4charts from "@amcharts/amcharts4/charts";
    import am4themes_animated from "@amcharts/amcharts4/themes/animated";
    

    const inputPkno = ref();
    const inputName = ref();
    const inputPrice = ref();
    const inputStock = ref();
    const inputIntroduce = ref();
    const inputInsertName = ref();
    const inputInsertPrice = ref();
    const inputInsertStock = ref();
    const inputInsertIntroduce = ref();
    const router = useRouter();
    const serachText = ref('');
    const num1 = ref(1);

    const activeName = ref('first')
    
    const dialogForImageUrl = ref('')
    const dialogVisibleImg = ref(false)
    const disabled = ref(false)
 
    function managerSignOut(){
        localStorage.clear()
        router.push({ name: "Product" });
    }
    const base64Img = ref([])
    function onfile(file){
        if(file.size/1024/10>=1){
            alert("檔案不可超過10KB")
            reload()
            return
        }
        const files = file.raw
        let reader = new FileReader();
        reader.readAsDataURL(files);
        reader.addEventListener("load",()=>{
            base64Img.value.push(reader.result);
        })
    }

    // beforeAvaterUpload:function()

    const handleClick = (tab: TabsPaneContext, event: Event) => {
        if(tab.index=="0"){
            //商品管理
            productManagement()
        } else if(tab.index == "1"){
            //訂單管理
            serachOrder()
        }else{
            //銷售報表
            let date = new Date();
            let nowMonth = String(date.getMonth()+1)
            if(nowMonth.length == 1){
                nowMonth = '0' + nowMonth
            }
            pieChart(nowMonth)
            ranking(nowMonth)
            barChart()
        }
    }

    let itemList = ref();
    let detailedItem = ref();
    let orderList = ref();
    let dialogFormVisible = ref(false);
    let dialogEditProduct = ref(false);
    let dialogInsertProduct = ref(false);
    let userInfo = localStorage.account;
    let jwtToken = localStorage.jwtToken;
    let permission = localStorage.permission;
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
        if(userInfo == null || jwtToken == null || permission != 1){
            router.push({ name: "Product" });
            return
        }
        productManagement()
    })
    function productSubmit(pkno){
        if(confirm('確定更改商品?')){
            axios
            .put('https://kinny-shop-backend-ezdtsdzl7q-uc.a.run.app/product/'+pkno, 
                {
                    name: inputName.value,
                    price : parseInt(inputPrice.value),
                    introduce:inputIntroduce.value,
                    stock:parseInt(inputStock.value)
                }, 
                {
                    headers:{"Authorization":`Bearer ${jwtToken}`}
                }
            )
            .then(function (response){
                alert(response.data.res)
                productManagement()
                reload()
            })
            .catch(function (error) {
            alert(error);
            });
        }else{
            return
        }
    }
    const isLock = ref(false);
    function insertProduct(){
        isLock.value = true;
        axios
        .post('https://kinny-shop-backend-ezdtsdzl7q-uc.a.run.app/product', 
            { 
                name : inputInsertName.value,
                price : parseInt(inputInsertPrice.value),
                stock : parseInt(inputInsertStock.value),
                introduce : inputInsertIntroduce.value,
                filePath: base64Img.value
            }, 
            {
                headers:{"Authorization":`Bearer ${jwtToken}`}
            }
        )
        .then(function (response){
            clearForm()
            dialogInsertProduct.value = false
            alert(response.data.res)
            reload()
            productManagement()
            isLock.value = false;
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
    function changeSendMark(pkno) {
        if(confirm('確定更改出貨狀態?')){
            axios
            .put('https://kinny-shop-backend-ezdtsdzl7q-uc.a.run.app/order/'+pkno, 
                {
                    
                }, 
                {
                    headers:{"Authorization":`Bearer ${jwtToken}`}
                }
            )
            .then(function (response){
                alert(response.data.res)
                serachOrder()
            })
            .catch(function (error) {
            alert(error);
            });
        }else{
            return
        }
    }
    const productList = ref();
    function productManagement(){
        axios
        .get('https://kinny-shop-backend-ezdtsdzl7q-uc.a.run.app/product', {headers:{"Authorization" : `Bearer ${jwtToken}`}} )
        .then(function (response){
            productList.value = response.data.res;
        })
        .catch(function (error) {
           alert(error);
        });
    }
    const productDetailedList = ref();
    function productDetailed(pkno){
        axios
        .get('https://kinny-shop-backend-ezdtsdzl7q-uc.a.run.app/product/detailed/'+pkno)
        //.get('http://localhost:80/detailed.php?pkno='+pkno)
        .then(function (response){
            productDetailedList.value = response.data.res;
            inputPkno.value = pkno;
            inputName.value = response.data.res.name;
            inputPrice.value = response.data.res.price;
            inputStock.value = response.data.res.stock;
            inputIntroduce.value = response.data.res.introduce;
        })
        .catch(function (error) {
           alert(error);
        });
    }
    function deleteProduct(pkno){
        axios
        .delete('https://kinny-shop-backend-ezdtsdzl7q-uc.a.run.app/product/'+pkno,{headers:{"Authorization" : `Bearer ${jwtToken}`}})
        //.get('http://localhost:80/detailed.php?pkno='+pkno)
        .then(function (response){
            if(confirm('確定要刪除?')){
                alert(response.data.res)
                productManagement()
                reload()
            }else{
                return
            }
        })
        .catch(function (error) {
           alert(error);
        });
    }
    const selectValue = ref()
    let tableData = ref()
    function ranking(month){
        axios
        .get('https://kinny-shop-backend-ezdtsdzl7q-uc.a.run.app/statistics/ranking/'+month, {headers:{"Authorization" : `Bearer ${jwtToken}`}} )
        .then(function (response){
            tableData.value = response.data.res;
        }).catch(function (error) {
           alert(error);
        });
    }
    function pieChart(month) {
        // pieChartGetValue();
        axios
        .get('https://kinny-shop-backend-ezdtsdzl7q-uc.a.run.app/statistics/'+month, {headers:{"Authorization" : `Bearer ${jwtToken}`}} )
        .then(function (response){
            tableData.value = response.data.res;
            console.log('tableData',tableData.value)
            // selectValue.value = response.data.res;
            // console.log('selectValue',selectValue.value )

            // Themes begin
            am4core.useTheme(am4themes_animated);
            // Themes end

            // Create chart instance
            const chart = am4core.create("chartdiv", am4charts.PieChart);
            // Add data
            chart.data = response.data.res;
            // Add and configure Series
            const pieSeries = chart.series.push(new am4charts.PieSeries());
            pieSeries.dataFields.value = "value";
            pieSeries.dataFields.category = "key";
            pieSeries.slices.template.stroke = am4core.color("#fff");
            pieSeries.slices.template.strokeOpacity = 1;

            // This creates initial animation
            pieSeries.hiddenState.properties.opacity = 1;
            pieSeries.hiddenState.properties.endAngle = -90;
            pieSeries.hiddenState.properties.startAngle = -90;

            chart.hiddenState.properties.radius = am4core.percent(0);
        
        })
        .catch(function (error) {
           alert(error);
        });
        
    }; // end am4core.ready()
    const selectChange = (selectValue) =>{
        pieChart(selectValue)
        ranking(selectValue)
    }
    function barChart() {
        axios
        .get('https://kinny-shop-backend-ezdtsdzl7q-uc.a.run.app/statistics', {headers:{"Authorization" : `Bearer ${jwtToken}`}} )
        .then(function (response){
            // Themes begin
            am4core.useTheme(am4themes_animated);
            // Themes end

            // Create chart instance
            let chart = am4core.create("barChart", am4charts.XYChart);
            chart.scrollbarX = new am4core.Scrollbar();

            // Add data
            chart.data = response.data.res;

           // Create axes
            var categoryAxis = chart.xAxes.push(new am4charts.CategoryAxis());
            categoryAxis.dataFields.category = "key";
            categoryAxis.renderer.grid.template.location = 0;
            categoryAxis.renderer.minGridDistance = 30;
            // categoryAxis.renderer.labels.template.horizontalCenter = "right";
            // categoryAxis.renderer.labels.template.verticalCenter = "middle";
            // categoryAxis.renderer.labels.template.rotation = 270;
            //categoryAxis.tooltip.disabled = true;
            categoryAxis.renderer.minHeight = 110;

            var valueAxis = chart.yAxes.push(new am4charts.ValueAxis());
            valueAxis.renderer.minWidth = 50;

            // Create series
            var series = chart.series.push(new am4charts.ColumnSeries());
            series.sequencedInterpolation = true;
            series.dataFields.valueY = "value";
            series.dataFields.categoryX = "key";
            series.tooltipText = "[{categoryX}: bold]{valueY}[/]";
            series.columns.template.strokeWidth = 0;

            //series.tooltip.pointerOrientation = "vertical";

            series.columns.template.column.cornerRadiusTopLeft = 10;
            series.columns.template.column.cornerRadiusTopRight = 10;
            series.columns.template.column.fillOpacity = 0.8;

            // on hover, make corner radiuses bigger
            var hoverState = series.columns.template.column.states.create("hover");
            hoverState.properties.cornerRadiusTopLeft = 0;
            hoverState.properties.cornerRadiusTopRight = 0;
            hoverState.properties.fillOpacity = 1;

            series.columns.template.adapter.add("fill", function(fill, target) {
                return chart.colors.getIndex(target.dataItem.index);
            });

            // Cursor
            chart.cursor = new am4charts.XYCursor();

            }); // end am4core.ready()
        
    }; // end am4core.ready()
    const selectOptions = [
        {
            value: '01',
            label: '一月',
        },
        {
            value: '02',
            label: '二月',
        },
        {
            value: '03',
            label: '三月',
        },
        {
            value: '04',
            label: '四月',
        },
        {
            value: '05',
            label: '五月',
        },
        {
            value: '06',
            label: '六月',
        },
        {
            value: '07',
            label: '七月',
        },
        {
            value: '08',
            label: '八月',
        },
        {
            value: '09',
            label: '九月',
        },
        {
            value: '10',
            label: '十月',
        },
        {
            value: '11',
            label: '十一月',
        },
        {
            value: '12',
            label: '十二月',
        },
    ]
    
    const handlePictureCardPreview = (file: UploadFile) => {
        dialogForImageUrl.value = file.url!
        dialogVisibleImg.value = true
    }

    const editNewProduct = ref(null);
    const upload = ref();
    function clearForm(){
        inputInsertName.value = null
        inputInsertPrice.value = null
        inputInsertStock.value = null
        inputInsertIntroduce.value = null
        base64Img.value = []
        // this.$refs.upload.clearForm
        console.log(upload)
        upload.value = null
    }
    function reload(){
        router.push({ name: "Tmp" });
    }
</script>

<style>
.managerSignOut{
    display: flex;
    justify-content: flex-end;
}
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
    width: 90%;
}
.tidyEditProduct p{
    margin-left: 140px;
    color: red;
}
.flexImg{
    width: 120px;
    height: 120px;
}
.editProductItem .el-upload{
    width: 120px;
    height: 120px;
}
.el-upload-list--picture-card .el-upload-list__item {
    width: 120px;
    height: 120px;
}
#chartdiv {
  width: 600px;
  height: 350px;
}
.productDistributed{
    font-family: Microsoft JhengHei;
    font-weight: bold;
    margin: 10px 0px 0px 5px;
    font-size: 20px;
}
.flex-table{
    width: 300px;
    height: 400px;
}
.flex-table .table{
    border-radius: 10px;
}
.flex-item{
    display: flex;
}
#barChart {
  width: 700px;
  height: 300px;
  margin: 15px 0;
}
</style>
<!-- <style scoped>
.avatar-uploader .avatar {
  width: 178px;
  height: 178px;
  display: block;
}
</style> -->