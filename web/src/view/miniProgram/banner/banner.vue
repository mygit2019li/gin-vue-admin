<template>
  <div>
    <div class="gva-table-box">
      <div class="gva-btn-list">
        <el-button type="primary" icon="plus" @click="openDialog">新增</el-button>
      </div>
      <el-table :data="tableData" style="width: 100%" row-key="ID">
        <el-table-column align="left" label="排序" prop="sort" width="80" />
        <el-table-column align="left" label="图片" prop="url" width="120">
          <template #default="scope">
            <CustomPic pic-type="file" :pic-src="scope.row.url" />
          </template>
        </el-table-column>
        <el-table-column align="left" label="名称" prop="name" width="120" />
        <el-table-column align="left" label="跳转链接" prop="link" min-width="150" />
        <el-table-column align="left" label="显示" prop="isShow" width="100">
          <template #default="scope">
            <el-switch
              v-model="scope.row.isShow"
              @change="toggleShow(scope.row)"
            />
          </template>
        </el-table-column>
        <el-table-column align="left" label="操作" width="160">
          <template #default="scope">
            <el-button type="primary" link icon="edit" @click="updateBannerAction(scope.row)">编辑</el-button>
            <el-popover v-model="scope.row.visible" placement="top" width="160">
              <p>确定要删除吗？</p>
              <div style="text-align: right; margin-top: 8px;">
                <el-button type="primary" link @click="scope.row.visible = false">取消</el-button>
                <el-button type="primary" @click="deleteBannerAction(scope.row)">确定</el-button>
              </div>
              <template #reference>
                <el-button type="primary" link icon="delete" @click="scope.row.visible = true">删除</el-button>
              </template>
            </el-popover>
          </template>
        </el-table-column>
      </el-table>
      <div class="gva-pagination">
        <el-pagination
          :current-page="page"
          :page-size="pageSize"
          :page-sizes="[10, 30, 50, 100]"
          :total="total"
          layout="total, sizes, prev, pager, next, jumper"
          @current-change="handleCurrentChange"
          @size-change="handleSizeChange"
        />
      </div>
    </div>
    <el-dialog v-model="dialogFormVisible" :before-close="closeDialog" :title="type==='create'?'新增轮播图':'编辑轮播图'">
      <el-form :model="form" label-width="80px">
        <el-form-item label="名称">
          <el-input v-model="form.name" autocomplete="off" />
        </el-form-item>
        <el-form-item label="图片">
          <div style="display: flex; align-items: center; gap: 10px;">
            <CustomPic v-if="form.url" pic-type="file" :pic-src="form.url" />
            <upload-image @on-success="(url) => form.url = url" />
          </div>
        </el-form-item>
        <el-form-item label="跳转链接">
          <el-input v-model="form.link" autocomplete="off" />
        </el-form-item>
        <el-form-item label="排序">
          <el-input-number v-model="form.sort" :min="0" />
        </el-form-item>
        <el-form-item label="显示">
          <el-switch v-model="form.isShow" />
        </el-form-item>
      </el-form>
      <template #footer>
        <div class="dialog-footer">
          <el-button @click="closeDialog">取 消</el-button>
          <el-button type="primary" @click="enterDialog">确 定</el-button>
        </div>
      </template>
    </el-dialog>
  </div>
</template>

<script setup>
import {
  createBanner,
  deleteBanner,
  updateBanner,
  getBannerList
} from '@/api/miniProgram'
import UploadImage from '@/components/upload/image.vue'
import CustomPic from '@/components/customPic/index.vue'
import { ref } from 'vue'
import { ElMessage } from 'element-plus'

defineOptions({
  name: 'BannerManage'
})

const form = ref({
  name: '',
  url: '',
  link: '',
  sort: 0,
  isShow: true
})

const page = ref(1)
const total = ref(0)
const pageSize = ref(10)
const tableData = ref([])

const handleSizeChange = (val) => {
  pageSize.value = val
  getTableData()
}

const handleCurrentChange = (val) => {
  page.value = val
  getTableData()
}

const getTableData = async() => {
  const table = await getBannerList({ page: page.value, pageSize: pageSize.value })
  if (table.code === 0) {
    tableData.value = table.data.list
    total.value = table.data.total
    page.value = table.data.page
    pageSize.value = table.data.pageSize
  }
}

getTableData()

const dialogFormVisible = ref(false)
const type = ref('')

const toggleShow = async(row) => {
  const res = await updateBanner(row)
  if (res.code === 0) {
    ElMessage.success('更新成功')
  }
}

const updateBannerAction = (row) => {
  type.value = 'update'
  form.value = { ...row }
  dialogFormVisible.value = true
}

const closeDialog = () => {
  dialogFormVisible.value = false
  form.value = {
    name: '',
    url: '',
    link: '',
    sort: 0,
    isShow: true
  }
}

const deleteBannerAction = async(row) => {
  row.visible = false
  const res = await deleteBanner({ ID: row.ID })
  if (res.code === 0) {
    ElMessage.success('删除成功')
    if (tableData.value.length === 1 && page.value > 1) {
      page.value--
    }
    getTableData()
  }
}

const enterDialog = async() => {
  let res
  if (type.value === 'create') {
    res = await createBanner(form.value)
  } else {
    res = await updateBanner(form.value)
  }

  if (res.code === 0) {
    ElMessage.success('保存成功')
    closeDialog()
    getTableData()
  }
}

const openDialog = () => {
  type.value = 'create'
  dialogFormVisible.value = true
}

</script>
