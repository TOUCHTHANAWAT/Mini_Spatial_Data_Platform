import { Layout, Menu } from "antd";
import { Link, Outlet } from "react-router-dom";

const { Header, Content } = Layout;

export default function MainLayout() {
  return (
    <Layout>
      <Header>
        <Menu theme="dark" mode="horizontal">
          <Menu.Item key="1">
            <Link to="/">Home</Link>
          </Menu.Item>
          <Menu.Item key="2">
            <Link to="/places">Places</Link>
          </Menu.Item>
          <Menu.Item key="3">
            <Link to="/places/new">Add</Link>
          </Menu.Item>
        </Menu>
      </Header>

      <Content style={{ padding: 15 }}>
        <Outlet />
      </Content>
    </Layout>
  );
}