import DeleteAllUsers from './components/DeleteAllUsers';
import GenerateUsers from './components/GenerateUsers';
import GetUser from './components/GetUser';
import GetUsers from './components/GetUsers';

function App() {
	return (
		<div>
			<GetUser />
			<GetUsers />
			<GenerateUsers />
			<DeleteAllUsers />
		</div>
	);
}

export default App;
