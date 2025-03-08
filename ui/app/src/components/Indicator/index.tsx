import './styles.css'

export function Indicator({
    total,
    current,
    onItemClick,
}: {
    total: number,
    current: number,
    onItemClick: (index: number) => void
}) {
    return (
        <view class='indicator'>
            {Array.from({ length: total }).map((_, idx) => (
                <IndicatorItem key={idx} active={idx === current} index={idx} onClick={onItemClick} />
            ))}
        </view>
    )
}

function IndicatorItem({
    active,
    index,
    onClick,
}: {
    active: boolean,
    index: number,
    onClick: (idx: number) => void
}) {
    return (
        <view
            class={`indicator-item ${active ? 'active' : ''}`}
            bindtap={() => onClick(index)}>
        </view>
    )
}