// import React from 'react';
// import { Card, CardContent, Typography, Button } from '@mui/material';

// interface PetitionProps {
//   title: string;
//   description: string;
//   link?: string;
//   signs: number;
// }

// const Petition: React.FC<PetitionProps> = ({ title, description, link, signs}) => {
//   return (
//     <Card sx={{ maxWidth: 600, margin: '20px auto', padding: '20px' }}>
//       <CardContent>
//         <Typography variant="h5" component="div" gutterBottom>
//           {title}
//         </Typography>
//         <Typography variant="body2" color="text.secondary" sx={{ marginBottom: '20px' }}>
//           {description}
//         </Typography>
//         {link ? (
//           <Button
//             variant="contained"
//             color="primary"
//             href={link}
//             target="_blank"
//             rel="noopener noreferrer"
//             sx={{ textTransform: 'none' }}
//           >
//             Read Full Petition
//           </Button>
//         ) : (
//           <Typography variant="body2" color="text.secondary" sx={{ marginTop: '10px' }}>
//             Link not available
//           </Typography>
//         )}
//       </CardContent>
//     </Card>
//   );
// };

// export default Petition;
