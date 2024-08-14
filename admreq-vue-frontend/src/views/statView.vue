<template>
    <v-row>
    <v-col cols="4">
    <v-card class="ma-3 pa-3" variant="flat">
        <p class="text-disabled font-weight-bold">Статистика по отделам:</p>
        
        <v-card variant="flat" class="ma-3" v-for="depart in StatStore.statData.byDepartment">
            <p>{{ depart.departmentName }}:</p> <p class="ml-3">В работе: <span class="text-red">{{ depart.tsrInWork }}</span>, Завершено: <span class="text-green">{{ depart.tsrFinished }}</span>, Принято: <span class="text-blue">{{ depart.tsrApplyed }}</span></p>
        </v-card>
    </v-card>
</v-col>
<v-col cols="4">
    <v-card class="ma-3 pa-3" variant="flat">
        <p class="text-disabled font-weight-bold">Статистика по исполнителям:</p>
        <v-card variant="flat" class="ma-3" v-for="employee in StatStore.statData.byEmployee">
            <p>{{ employee.employeeName }}:</p> <p class="ml-3">В работе: <span class="text-red">{{ employee.tsrInWork }}</span>, Завершено: <span class="text-green">{{ employee.tsrFinished }}</span>, Принято: <span class="text-blue">{{ employee.tsrApplyed }}</span></p>
        </v-card>
    </v-card>
</v-col>
    <!-- <span>{{ StatStore.statData }}</span> -->
</v-row>
</template>

<script setup>
import { useStatStore } from '../stores/StatStore';
import { useAuthStore } from '../stores/AuthStore';
const StatStore = useStatStore();
const AuthStore = useAuthStore();

StatStore.getStatData(AuthStore.credentials.departmentId, AuthStore.credentials.token)
</script>