<template>
    <b-container class="home">
        <b-row class="justify-content-center">
            <b-col lg="4">
                <b-card header="Enter a domain" header-text-variant="light" header-bg-variant="dark" header-class="card-title" class="text-center">
                    <b-form @submit="onSubmit">
                        <b-form-input
                            v-model="newDomain"
                            required
                            placeholder="ej: google.com">
                        </b-form-input>
                        <b-button block type="submit" variant="danger">Submit</b-button>
                    </b-form>
                </b-card>
            </b-col>
            <b-col lg="8">
                <b-card header="Dominio" header-text-variant="light" header-bg-variant="dark" header-class="card-title">
                    <h3><strong>{{domain.title}}</strong></h3>
                    <h3 v-if="domain.title == null"><strong>Waiting for a domain</strong></h3>
                    <h3 v-if="domain.title == ''"><strong>Try later with this domain.</strong></h3>
                    <div class="info-domain">
                        <p><strong>Ssl grade:</strong>{{domain.ssl_grade}}</p>
                        <p><strong>Previous ssl grade:</strong>{{domain.previous_ssl_grade}}</p>
                        <p><strong>Is down:</strong>{{domain.is_down}}</p>
                        <p><strong>Servers changed:</strong>{{domain.servers_changed}}</p>
                    </div>
                    <h4>Servers</h4>
                    <b-table class="text-center" striped :items="domain.servers"></b-table>
                </b-card>
            </b-col>
        </b-row>        
    </b-container>
</template>

<script>
import axios from "axios"

export default {
    data () {
        return {
            newDomain: "",
            domain: {

            }
        }
    },
    methods: {
        onSubmit(e) {
            e.preventDefault()  

            if (this.newDomain != "") {

                axios.get("api/servers/search/" + this.newDomain)
                .then(response => {
                    this.domain = response.data
                })
                .catch(err => console.error(err))
            }
        }
    }
}
</script>

<style>
.card-title{
    font-size: 25px;
}

.home{
    margin-top: 40px;
}

form input, form button {
    margin-top: 15px;
    font-size: 20px;
}

.info-domain{
    grid-template-columns: 1fr 1fr;
}
</style>