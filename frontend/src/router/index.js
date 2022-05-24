import { createRouter, createWebHistory } from "vue-router";
import Body from '../components/BodyComponent.vue'
import Criar from '../components/CriadorPersonagens.vue'
import Instrucoes from '../components/ComoJogar.vue'
import Lovecraft from '../components/SobreLovecraft.vue'
import Comprar from '../components/ComprarProdutos.vue'
import Fonte from '../components/CodigoFonte.vue'

const routes = [ 
	{
		path: '/',
		name: 'Home',
		component: Body,
	},
	{
		path: '/criar',
		name: 'Criar',
		component: Criar
	},
	{
		path: '/instrucoes',
		name: 'Instruções',
		component: Instrucoes 
	},
	{
		path: '/lovecraft',
		name: 'Lovecraft',
		component: Lovecraft 
	},
	{
		path: '/comprar',
		name: 'Comprar',
		component: Comprar 
	},
	{
		path: '/fonte',
		name: 'Fonte',
		component: Fonte 
	}
]

const router = createRouter({ history: createWebHistory(), routes })
export default router
