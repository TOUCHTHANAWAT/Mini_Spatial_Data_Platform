import { Layout, Menu } from "antd";
import { Link, Outlet } from "react-router-dom";
import "./MainLayout.css";

const { Header, Content } = Layout;

export default function MainLayout() {
  return (
    <Layout>
      <Header className="nav-bar">
        <div className="nav-left">
          <img src="/geographical.png" alt="logo" className="logo-img" />
          <div className="logo">Mini Spatial Data</div>

          <Menu mode="horizontal" className="menu">
            <Menu.Item key="1">
              <Link to="/">Map View</Link>
            </Menu.Item>
            <Menu.Item key="2">
              <Link to="/places">Table View</Link>
            </Menu.Item>
            {/* <Menu.Item key="3">
              <Link to="/places/new">Add</Link>
            </Menu.Item> */}
          </Menu>
        </div>
      </Header>

      <Content className="content">
        <Outlet />
      </Content>
    </Layout>
  );
}