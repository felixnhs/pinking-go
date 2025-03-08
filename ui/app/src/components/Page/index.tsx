import type { CSSProperties } from "@lynx-js/types";
import type { ReactNode } from "react";
import './styles.css';
import SafeArea from "../SafeArea/index.jsx";
import NavBar from "../NavBar/index.jsx";
import favoriteIcon from '../../assets/heart.png';

export function Page({
    children,
    title,
    style
}: {
    children: ReactNode,
    title?: string,
    style?: CSSProperties
}) {
    return (
        <SafeArea style={style}>
            <view class="page-container">
                <NavBar title={title} />
                {children}
                <view class="card-detail-container">
                    <view class="card-detail">
                        <view class="card-detail-title">
                            <text class="card-detail-title-text">Lorem Ipsum</text>
                            <image mode="aspectFill" class="favorite-icon" src={favoriteIcon} style={{ width: '20px', height: '20px'}} /> 
                        </view>
                        <view class="card-detail-desc">
                            <text class="card-detail-desc-autor">Felix</text>
                            <text class="card-detaul-desc-date">3 days ago</text>
                        </view>
                    </view>
                </view>
            </view>
        </SafeArea>
    )
}