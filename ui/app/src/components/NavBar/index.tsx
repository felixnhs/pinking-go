import './styles.css'
import  backIcon from '../../assets/back.png';
import  starIcon from '../../assets/star.png';

interface NavBarProps {
    onBack?: () => void,
    title?: string
    onFavorite?: () => void
}

export default function NavBar({ onBack, title, onFavorite }: NavBarProps) {
    return (
        <view className="nav-bar">
            <image className="left-icon" src={backIcon} bindtap={onBack} />
            <text className="nav-title">{title}</text>
            <image className="right-icon" src={starIcon} bindtap={onFavorite} />
        </view>
    )
}