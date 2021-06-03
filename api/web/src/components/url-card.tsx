
import "./url-card.css"
import {Divider, Card, message, Tooltip} from "antd";
import { UrlResult } from "../models/urls";
import { CopyOutlined } from '@ant-design/icons';

type Props = {
    originalUrl: string,
    shortUrl: string
}



export function UrlCard({ originalUrl, shortUrl} : UrlResult){

    return(
        <Card title="Your Most Recent Url">
            <div className="url-card-container">
                <Card className="url-card-container-panel">
                    <a target="_blank" rel="noreferrer"  href={originalUrl}>
                        {originalUrl}
                    </a>
                </Card>
                <Card className="url-card-container-panel">
                    <a target="_blank" rel="noreferrer"  href={shortUrl}>
                        { shortUrl }
                    </a>
                    <Tooltip title={"Copy to Clipboard"}>
                        <CopyOutlined style={{ position: "absolute", right: 5, bottom: 5, color: "#959292"}}
                                      size={40}
                                      onClick={async () => {
                                          await navigator.clipboard.writeText(shortUrl)
                                          showCopyToClipboardMessage()}
                                      }
                        />
                    </Tooltip>
                </Card>
            </div>
        </Card>
    )
}

const showCopyToClipboardMessage = () => {
    message.success('Copied to Clipboard!');
};
