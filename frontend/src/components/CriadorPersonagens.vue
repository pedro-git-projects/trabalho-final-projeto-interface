<template>
<div class="wrapper d-flex flex-column min-vh-100">
	<div class="container">
		<div class="row">
			<div class="col">
				<h1 class="mt-5">Gerador de Investigadores</h1>
				<hr>
				<form-tag @myevent="submitHandler" name="myform" :event="myevent">
				<text-input
					v-model="nome"
					label="Nome"
					type="text"
					name="nome"
					required="true"
				>
				</text-input>

				<text-input
					v-model="idade"
					label="Idade"
					type="number"
					name="idade"
					required="true"
				>
				</text-input>

				<text-input
					v-model="residencia"
					label="Residência"
					type="text"
					name="residencia"
					required="true"
				>
				</text-input>

				<text-input
					v-model="residencia"
					label="Local de Nascimento"
					type="text"
					name="local de nascimento"
					required="true"
				>
				</text-input>

				<text-input
					v-model="ocupacao"
					label="Ocupação"
					type="text"
					name="ocupacao"
					required="true"
				>
				</text-input>

				<hr>
				
				<input type="submit" class="btn btn-primary" value="criar">
				</form-tag>
			</div>
		</div>
	</div>
</div>
</template>

<script>
import TextInput from './TextInput.vue'
import FormTag from './FormTag.vue'
export default {
	name: 'criar-investigador',
	components: {
		FormTag,
		TextInput,
	},
	data() {
		return {
			nome: "",
			idade: "",
			residencia: "",
			nascimento: "",
			ocupacao: "",
		}	
	},
	methods: {
		submitHandler() {
			console.log("Metodo invocado")
		}
	},
}
</script>
