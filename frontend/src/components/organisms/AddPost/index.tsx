'use client'
import React, { useState } from 'react'
import './index.scss'
import PersonIcon from '@mui/icons-material/Person';
import { WButtonPost } from '@/components';
import ImageIcon from '@mui/icons-material/Image';
import GifBoxIcon from '@mui/icons-material/GifBox';
import SignalCellularAltIcon from '@mui/icons-material/SignalCellularAlt';
import SentimentSatisfiedAltIcon from '@mui/icons-material/SentimentSatisfiedAlt';
import CalendarTodayIcon from '@mui/icons-material/CalendarToday';
import AddCircleOutlineIcon from '@mui/icons-material/AddCircleOutline';
import PublicIcon from '@mui/icons-material/Public';
import WTextArea from '@/components/atoms/TextArea/textArea';

interface AddPostProps {
    imageUrl: string
}

const AddPost: React.FC<AddPostProps> = ({imageUrl}) => {

    const [textAreaValue, setTextAreaValue] = useState("Holaa");
    
    const handleChangeTextArea = (event: any) => {
        setTextAreaValue(event.target.value);
    }

    return(
        <div className="AddPost_main_container">
            <div className="avatar_container">
                {imageUrl ? <img src={imageUrl} /> : <PersonIcon fontSize='large'/>}
                
            </div>
            <div className="post_container">
                <div className="text_area_container">
                    <WTextArea placeholder='Escribe...' textAreaValue={textAreaValue} handleChangeTextArea={handleChangeTextArea}/>
                </div>
                <div>
                    <a href="#"><PublicIcon/>  Everyone can reply</a>
                    <br />
                    <hr />
                </div>
                <div className="post_buttons_container">
                    <div className="buttons_container_one">
                        <button><ImageIcon fontSize='large'/></button>
                        <button><GifBoxIcon fontSize='large'/></button>
                        <button><SignalCellularAltIcon fontSize='large'/></button>
                        <button><SentimentSatisfiedAltIcon fontSize='large'/></button>
                        <button><CalendarTodayIcon fontSize='large'/></button>
                    </div>
                    <div className="buttons_container_two">
                        <button><AddCircleOutlineIcon fontSize='large'/></button>
                        <WButtonPost 
                            typeColor='primary'
                            text='Post'
                            disabled={false}
                        />
                    </div>
                </div>
            </div>
        </div>
    );
};

export default AddPost;

AddPost.defaultProps = {
    imageUrl: './images/Ellipse 91.svg'
}