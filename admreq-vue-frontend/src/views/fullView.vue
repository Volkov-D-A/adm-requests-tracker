<template>
    <v-row>
        <v-col cols="3">
            <div class="ma-3">
                <p class="text-disabled font-weight-bold">Автор обращения:</p>
                <p class="ml-5 mb-2">{{ FullStore.fullTicket.userLastname }}&nbsp;{{ FullStore.fullTicket.userFirstname }}&nbsp;{{ FullStore.fullTicket.userSurname }}</p>
                <p class="text-disabled font-weight-bold">Исполнитель обращения:</p>
                <p class="ml-5 mb-2">{{ FullStore.fullTicket.employeeLastname }}&nbsp;{{ FullStore.fullTicket.employeeFirstname }}&nbsp;{{ FullStore.fullTicket.employeeSurname }}</p>
                <v-form v-if="AuthStore.credentials.Role === 'admin' && mode === 'admin'" fast-fail @submit.prevent="FullStore.setEmployee(FullStore.fullTicket.id, employeeId, AuthStore.credentials.token)">
                    <v-select v-model="employeeId" :items=employees></v-select>
                    <v-btn type="submit" color="primary" block class="mt-2">Назначить</v-btn>
                </v-form>
                <p><span class="text-disabled font-weight-bold">Важность обращения:</span>&nbsp;<span v-if="FullStore.fullTicket.important" class="text-red">Высокая</span><span v-if="!FullStore.fullTicket.important" class="text-green">Обычная</span>
                   <v-switch v-if="AuthStore.credentials.Role === 'admin' && mode === 'admin'"
                   v-model="FullStore.fullTicket.important"
                   @change="FullStore.toggleImportance(FullStore.fullTicket.id, AuthStore.credentials.token, FullStore.fullTicket.important)"
                   >
                    </v-switch>
                </p>
                <p class="text-disabled font-weight-bold">Дата обращения:</p>
                <p class="ml-5 mb-2">{{ AuthStore.myDateTimeFormat(FullStore.fullTicket.postedAt) }}</p>
                <p class="text-disabled font-weight-bold" v-if="FullStore.fullTicket.finished === true">Обращение завершено: </p>
                <p v-if="FullStore.fullTicket.finished === true" class="ml-5 mb-2">{{ AuthStore.myDateTimeFormat(FullStore.fullTicket.finishedAt) }}</p>
                <div>
                <v-form v-if="AuthStore.credentials.Role != 'user' && FullStore.fullTicket.finished === false && mode === 'employee'" fast-fail @submit.prevent="FullStore.finishTstr(FullStore.fullTicket.id, AuthStore.credentials.token)">
                    <v-btn type="submit" color="red" block class="mt-2">Завершить</v-btn>
                </v-form>
                <v-form v-if="FullStore.fullTicket.finished === true && FullStore.fullTicket.applied === false && mode === 'user'" fast-fail @submit.prevent="FullStore.applyTstr(FullStore.fullTicket.id, AuthStore.credentials.token)">
                    <v-btn type="submit" color="blue" block class="mt-2">Принять</v-btn>
                </v-form>
                </div>
            </div>    
        </v-col>
        <v-col cols="3">
            <div class="ma-3">
                <p class="text-disabled font-weight-bold">Текст обращения:</p>
                <p class="ml-5">{{ FullStore.fullTicket.text }}</p>
            </div>
        </v-col>
        <v-col cols="3">
            <v-card class="pa-2 mt-2" variant="flat">
                <p class="text-disabled font-weight-bold">Сообщения:</p>
                <v-card
                        v-if="FullStore.count > 0"
                        v-for="comment in FullStore.comments"
                        :key="comment.commId"
                        :text=comment.CommentText
                        :subtitle='comment.lastname + " " + comment.firstname[0] + "." + comment.surname[0] + ". // " +  AuthStore.myDateTimeFormat(comment.postedAt)'
                        variant="flat"
                ></v-card>
                <v-card 
                    v-if="FullStore.count === 0"
                    variant="flat"
                    >
                    <span class="ma-2">Сообщений нет!</span><br><br>
                </v-card>
            </v-card>
        </v-col>
        <v-col>
            <v-card class="pa-2 mr-2 mt-2" variant="elevated" elevation="16" color="teal-lighten-4">
                <span class="ma-2">Написать сообщение:</span><br><br>
                <v-form fast-fail @submit.prevent="FullStore.sendTicketComment(AuthStore.credentials.token, id, message)">
                    <v-textarea counter auto-grow variant="outlined" background-color="blue-lighten-5" v-model="message" label="Текст сообщения"></v-textarea>
                    <v-btn type="submit" color="teal-darken-1" block class="mt-2">Отправить</v-btn>
                </v-form>
            </v-card>
            <div class="pa-3">
            </div>
        </v-col>
    </v-row>
</template>

<script setup>
import { useAuthStore } from '../stores/AuthStore';
import { useFullStore } from '../stores/FullStore';
import { useUsersStore } from '../stores/UsersStore';
import { useRoute } from 'vue-router';
import { ref } from 'vue';
const AuthStore = useAuthStore();
const FullStore =  useFullStore();
const UsersStore = useUsersStore();
const route = useRoute();
const id = route.params.id
const mode = route.params.mode
UsersStore.getUsers(AuthStore.credentials.token)
FullStore.getFullTicket(AuthStore.credentials.token, id)
FullStore.getTicketComments(AuthStore.credentials.token, id)
const employees = UsersStore.getEmployeeItems(AuthStore.credentials.token)
const message = ref("")
const employeeId = ref("")
</script>