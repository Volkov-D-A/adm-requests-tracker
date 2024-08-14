<template>
    <v-col cols="6">
    <v-card class="ma-3 pa-3">
        <p class="text-disabled font-weight-bold">Статистика по отделам:</p>
        
        <v-card variant="flat" class="ma-3" v-for="depart in StatStore.statData">
            <p>{{ depart.departmentName }}:</p> <p class="ml-3">В работе: <span class="text-red">{{ depart.tsrInWork }}</span>, Завершено: <span class="text-green">{{ depart.tsrFinished }}</span>, Принято: <span class="text-blue">{{ depart.tsrApplyed }}</span></p>
        </v-card>
    </v-card>
</v-col>
    <!-- <span>{{ StatStore.statData }}</span> -->
</template>

<script setup>
import { useStatStore } from '../stores/StatStore';
import { useAuthStore } from '../stores/AuthStore';
const StatStore = useStatStore();
const AuthStore = useAuthStore();

StatStore.getStatData(AuthStore.credentials.departmentId, AuthStore.credentials.token)
</script>