import React from 'react';
import { Layout } from 'antd';
import ContainerTable from './components/ContainerTable';

const { Header, Content } = Layout;

const App: React.FC = () => {
    return (
        <Layout>
            <Header>
                <h1 style={{ color: 'white' }}>Container Monitoring</h1>
            </Header>
            <Content style={{ padding: '50px' }}>
                <ContainerTable />
            </Content>
        </Layout>
    );
};

export default App;