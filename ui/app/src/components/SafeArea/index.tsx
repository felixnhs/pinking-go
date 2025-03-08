import type { CSSProperties } from "@lynx-js/types";
import type { ReactNode } from "react";
import './styles.css'

export default function SafeArea({
    children,
    style
}: {
    children: ReactNode,
    style?: CSSProperties
}) {
    const isIOS = SystemInfo.platform === 'iOS';

    return (
        <view class={`safe-area ${isIOS ? 'ios' : 'android'}`} style={style}>
            {children}
        </view>
    )
}