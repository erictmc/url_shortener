import React, { useState } from 'react';
import './App.css';
import 'antd/dist/antd.css';
import { Layout, Input, Button, Col, Row, message, Divider, Card } from 'antd';
import {isValidUrl} from "./utils";
import * as urlApi from "./api/urls-api"
import { UrlResult } from "./models/urls";
import {UrlCard} from "./components/url-card";

const { Header, Content, Footer } = Layout;


export function App() {
    const [url, setUrl] = useState("");
    const [createdUrl, setCreatedUrl ] = useState<UrlResult | undefined >(undefined)
    const [apiCallInProgress, setApiCallInProgress] = useState(false);

    const sanitizeUrl = (rawUrl: string) => {
        if (rawUrl.includes("http://") || rawUrl.includes("https://")){
            return rawUrl
        }
        return "https://" + rawUrl
    }

    const checkUrl = () => {
        return isValidUrl(sanitizeUrl(url));
    };

    const createUrl = async () => {
        try {
            setApiCallInProgress(true)
            const urlResult = await urlApi.createUrl({ originalUrl: sanitizeUrl(url) })
            success()
            setUrl("")
            setCreatedUrl(urlResult)
        } catch(e){
            error()
        } finally {
            setApiCallInProgress(false)
        }
    }

  return (
      <Layout className="layout" >
          <Header>
              <div className="logo" />
          </Header>
          <Content style={{ padding: '0 50px' }} className="site-layout-content">
              <Row>
                  <Col span={12} offset={6}>
                      <div id="form-container">
                          <Input addonBefore="https://"
                                 value={url}
                                 placeholder={"www.example.com"}
                                 onChange={(event) => setUrl(event.target.value)}
                          />
                          <div style={{width: "20px"}}/>
                          <Button type="primary"
                                  shape="round"
                                  disabled={!checkUrl()}
                                  loading={apiCallInProgress}
                                  onClick={createUrl}
                          >
                              Create Tiny Url
                          </Button>
                      </div>
                  </Col>
              </Row>
              { createdUrl && (
                  <>
                      <Divider dashed={true}/>
                          <Row>
                          <Col span={12} offset={6}>
                              <UrlCard originalUrl={createdUrl.originalUrl}
                                       shortUrl={createdUrl.shortUrl}
                              />
                          </Col>
                      </Row>
                  </>
                  )
              }
          </Content>
          <Footer style={{ textAlign: 'center' }}>
              Url Shortener ©2021
          </Footer>
      </Layout>
  );
}


const success = () => {
    message.success('Your url has been created!');
};

const error = () => {
    message.error('There was an error when creating your url!');
};
